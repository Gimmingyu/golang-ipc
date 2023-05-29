package udp

import (
	"net"
)

func Listen(port int) (*net.UDPConn, error) {
	return net.ListenUDP("udp", &net.UDPAddr{Port: port})
}

func Read(conn *net.UDPConn, callback func(*net.UDPConn, []byte, net.Addr)) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}
		go callback(conn, buf, addr)
	}
}
