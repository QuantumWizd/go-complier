package api

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/QuantumWizd/go-complier/chat"
	"github.com/QuantumWizd/go-complier/network"
)

func RunHost() {
	fs := flag.NewFlagSet("host", flag.ExitOnError)
	name := fs.String("name", "", "Your display name")
	port := fs.Int("port", 8888, "Port to listen on")
	fs.Parse(os.Args[2:])

	if *name == "" {
		*name = promptName()
	}

	ip, err := network.GetPrivateIP()
	if err != nil {
		chat.PrintError("Could not detect private IP: " + err.Error())
		os.Exit(1)
	}

	roomCode := generateRoomCode()

	chat.ClearScreen()
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════╗")
	fmt.Println("  ║         Room Created!            ║")
	fmt.Printf("  ║   Room Code : %-18s║\n", roomCode)
	fmt.Printf("  ║   Your IP   : %-18s║\n", ip)
	fmt.Printf("  ║   Host      : %-18s║\n", *name)
	fmt.Println("  ║                                  ║")
	fmt.Println("  ║  Share the room code with your   ║")
	fmt.Println("  ║  teammates on the same network   ║")
	fmt.Println("  ╚══════════════════════════════════╝")
	fmt.Println()

	// start UDP broadcaster
	stopBroadcast := make(chan struct{})
	go network.StartBroadcast(roomCode, ip, stopBroadcast)

	// start TCP listener
	addr := fmt.Sprintf(":%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		chat.PrintError(fmt.Sprintf("Port %d is already in use. Is another chat session running?", *port))
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("  Listening on %s:%d — waiting for users...\n\n", ip, *port)

	room := chat.NewRoom(ip)

	// handle Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println("\n\n  Host is closing the room...")
		room.Broadcast(chat.Message{
			Type: chat.TypeSys,
			From: "HOST",
			Body: "room_closed",
		}, "")
		close(stopBroadcast)
		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()

	// accept connections loop
	userCounter := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		userCounter++
		userID := fmt.Sprintf("user-%d", userCounter)
		go room.HandleUser(conn, userID)
	}
}

func generateRoomCode() string {
	rand.Seed(time.Now().UnixNano())
	chars := "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	code := make([]byte, 4)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func promptName() string {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	return name
}
