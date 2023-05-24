package ipc

import (
	"log"
	"net"
	"net/rpc"
)

type InterProcessCallListener interface {
	Listen()
}

func RegisterServer(server InterProcessCallListener) error {
	return rpc.Register(server)
}

func Listen(listener net.Listener) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("recovered from panic : %v", r)
				}
			}()

			for {
				conn, err := listener.Accept()
				if err != nil {
					continue
				}

				go rpc.ServeConn(conn)
			}
		}()
	}
}