package main

import (
	"golang-ipc/pkg/ipc"
	"golang-ipc/pkg/types"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
}

func (s *Server) PrintOut(args *types.Args, result *types.Resp) error {

	log.Println(args.Msg)

	result.Success = true
	return nil
}

func (s *Server) Listen() {
	ipc.Listen(s.listener)
}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}

	server := &Server{listener: listener}

	if err := ipc.RegisterServer(server); err != nil {
		log.Fatal(err)
	}

	server.Listen()
}