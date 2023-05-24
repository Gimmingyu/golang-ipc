package ipc

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type InterProcessCallDialer interface {
	Dial()
	Call(method string, args interface{}, result interface{}) error
}

func Dial() *rpc.Client {
	addr := fmt.Sprintf("localhost:%v", os.Getenv("IPC_PORT"))
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("failed to dial : %v", addr)
	}

	return client
}

func Call(client *rpc.Client, method string, args interface{}, result interface{}) error {
	if err := client.Call(method, args, result); err != nil {
		return err
	}
	return nil
}