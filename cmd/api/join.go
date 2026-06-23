package api

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/QuantumWizd/go-complier/chat"
	"github.com/QuantumWizd/go-complier/network"
)

func RunJoin() {
	fs := flag.NewFlagSet("join", flag.ExitOnError)
	name := fs.String("name", "", "Your display name")
	code := fs.String("code", "", "Room code to join")
	port := fs.Int("port", 8888, "Port to connect on")
	fs.Parse(os.Args[2:])

	if *name == "" {
		*name = promptName()
	}

	if *code == "" {
		fmt.Print("Enter room code: ")
		fmt.Scanln(code)
	}

	*code = strings.ToUpper(strings.TrimSpace(*code))

	if len(*code) != 4 {
		chat.PrintError("Room code must be 4 characters")
		os.Exit(1)
	}

	fmt.Printf("\n  Searching for room %s on your network...\n", *code)

	hostIP, err := network.DiscoverHost(*code)
	if err != nil {
		chat.PrintError(err.Error())
		os.Exit(1)
	}

	fmt.Printf("  Found room at %s — connecting...\n", hostIP)

	addr := fmt.Sprintf("%s:%d", hostIP, *port)

	var conn net.Conn
	for attempt := 1; attempt <= 15; attempt++ {
		conn, err = net.Dial("tcp", addr)
		if err == nil {
			break
		}
		fmt.Printf("  Attempt %d/15 failed, retrying...\n", attempt)
		waitWithSleep(2)
	}

	if conn == nil {
		chat.PrintError(fmt.Sprintf("Could not connect to room %s after 15 attempts", *code))
		os.Exit(1)
	}

	fmt.Printf("  Connected to room %s!\n\n", *code)

	// send join message
	joinMsg := chat.Message{
		Type: chat.TypeSys,
		From: *name,
		Body: "joined",
	}
	conn.Write([]byte(chat.Format(joinMsg)))

	// start chat session
	chat.Start(conn, *name)
}

func waitWithSleep(seconds int) {
	for i := 0; i < seconds; i++ {
		fmt.Print(".")
		waitOneSec()
	}
	fmt.Println()
}

func waitOneSec() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		c := make(chan struct{})
		go func() {
			defer close(c)
		}()
		<-c
	}()
	<-done
}
