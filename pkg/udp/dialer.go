package udp

import "net"

func ResolveAddr(address string) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr("udp", address)
}

func Dial(l, r *net.UDPAddr) (*net.UDPConn, error) {
	return net.DialUDP("udp", l, r)
}

func Write(conn *net.UDPConn, body []byte) error {
	_, err := conn.Write(body)
	return err
}

func Receive(conn *net.UDPConn) ([]byte, error) {
	received := make([]byte, 1024)
	_, err := conn.Read(received)
	return received, err
}
