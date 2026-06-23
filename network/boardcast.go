package network

import (
	"fmt"
	"net"
	"time"
)

func StartBroadcast(roomCode string, hostIP string, stop chan struct{}) {
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9999")
	if err != nil {
		fmt.Println("Broadcast error:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Broadcast dial error:", err)
		return
	}
	defer conn.Close()

	message := fmt.Sprintf("ROOM:%s:%s", roomCode, hostIP)

	for {
		select {
		case <-stop:
			return
		default:
			conn.Write([]byte(message))
			time.Sleep(2 * time.Second)
		}
	}
}