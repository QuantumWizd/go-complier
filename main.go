package main

import (
	"fmt"
	"os"

	"github.com/QuantumWizd/go-complier/cmd/api"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "host":
		api.RunHost()
	case "join":
		api.RunJoin()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`
  YourNet — CLI Group Chat over LAN

  Usage:
    chat host [--name <yourname>] [--port <port>]
    chat join [--name <yourname>] [--code <roomcode>] [--port <port>]

  Examples:
    chat host --name alice
    chat join --name bob --code KJ7X
    chat join --code KJ7X         (will prompt for name)

  Commands:
    host    Create a new chat room and get a room code
    join    Join an existing room using the room code

  Flags:
    --name  Your display name (prompted if not provided)
    --code  4-character room code (join only)
    --port  Port number (default: 8888)
	`)
}
