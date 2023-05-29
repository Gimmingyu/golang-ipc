package udp

import "net"

type Listener interface {
	Read()
}

type Dialer interface {
	Dial(l, r *net.UDPAddr) (*net.UDPConn, error)
	Write(body []byte) error
	Receive() ([]byte, error)
}
