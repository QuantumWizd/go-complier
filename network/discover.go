package network

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func DiscoverHost(roomCode string) (string, error) {
	addr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		return "", fmt.Errorf("could not resolve UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return "", fmt.Errorf("could not listen on UDP port 9999: %w", err)
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(15 * time.Second))

	buf := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			return "", fmt.Errorf("room %s not found on your network", roomCode)
		}

		message := string(buf[:n])
		parts := strings.Split(message, ":")

		if len(parts) != 3 {
			continue
		}

		if parts[0] == "ROOM" && parts[1] == roomCode {
			return parts[2], nil
		}
	}
}