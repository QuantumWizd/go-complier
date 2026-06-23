package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
)

func Start(conn net.Conn, myName string) {
	var wg sync.WaitGroup
	quit := make(chan struct{})

	reader := bufio.NewReader(conn)
	var peerRoomCode string

	// read WHO and HIST before starting live session
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			PrintError("Lost connection during setup")
			return
		}

		msg, err := Parse(line)
		if err != nil {
			continue
		}

		if msg.Type == TypeWho {
			peerRoomCode = msg.Body
			users := strings.Split(msg.Body, ",")
			ClearScreen()
			PrintHeader("----", myName, "connecting...", len(users))
			PrintSystem("Connected to room")
			PrintUserList(users)
			continue
		}

		if msg.Type == TypeHist {
			PrintHistory(msg)
			continue
		}

		// first non-WHO non-HIST message means live chat has begun
		// process it then break
		if msg.Type == TypeMsg {
			PrintMessage(msg)
		}
		break
	}

	_ = peerRoomCode

	PrintSystem("Live chat started — type your message and press Enter")
	PrintPrompt()

	// goroutine 1 — reader
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				return
			default:
				line, err := reader.ReadString('\n')
				if err != nil {
					select {
					case <-quit:
					default:
						PrintSystem("Connection lost")
						close(quit)
					}
					return
				}

				msg, err := Parse(line)
				if err != nil {
					continue
				}

				switch msg.Type {
				case TypeMsg:
					fmt.Print("\r\033[2K")
					PrintMessage(msg)
				case TypeSys:
					fmt.Print("\r\033[2K")
					if msg.Body == "joined" {
						PrintSystem(fmt.Sprintf("%s has joined", msg.From))
					} else if msg.Body == "left" {
						PrintSystem(fmt.Sprintf("%s has left", msg.From))
					} else if msg.Body == "room_closed" {
						PrintSystem("Host has closed the room")
						close(quit)
						return
					}
				}
			}
		}
	}()

	// goroutine 2 — writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for {
			select {
			case <-quit:
				return
			default:
				if !scanner.Scan() {
					return
				}
				text := strings.TrimSpace(scanner.Text())
				if text == "" {
					PrintPrompt()
					continue
				}

				if text == "/users" {
					PrintPrompt()
					continue
				}

				if text == "/quit" {
					handleQuit(conn, myName, quit)
					return
				}

				msg := Message{Type: TypeMsg, From: myName, Body: text}
				_, err := conn.Write([]byte(Format(msg)))
				if err != nil {
					PrintError("Could not send message")
					return
				}
				PrintPrompt()
			}
		}
	}()

	// goroutine 3 — signal handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		for {
			select {
			case <-quit:
				return
			case <-sigChan:
				handleQuit(conn, myName, quit)
				return
			}
		}
	}()

	wg.Wait()
	conn.Close()
}

func handleQuit(conn net.Conn, myName string, quit chan struct{}) {
	fmt.Print("\n^C  Really quit? [y/N]: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer := strings.TrimSpace(scanner.Text())
		if answer == "y" || answer == "Y" {
			msg := Message{Type: TypeSys, From: myName, Body: "left"}
			conn.Write([]byte(Format(msg)))
			select {
			case <-quit:
			default:
				close(quit)
			}
			os.Exit(0)
		}
	}
	fmt.Print("\r\033[2K")
	PrintPrompt()
}