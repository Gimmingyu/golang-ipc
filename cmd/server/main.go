package main

import (
	"fmt"
	"golang-ipc/pkg/udp"
	"log"
	"net"
	"time"
)

func main() {
	server, err := udp.Listen(5050)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	for {
		buf := make([]byte, 1024)
		_, addr, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}

		go UdpResp(buf, server, addr)
	}
}

func UdpResp(buf []byte, server *net.UDPConn, addr net.Addr) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))

	server.WriteTo([]byte(responseStr), addr)
}
