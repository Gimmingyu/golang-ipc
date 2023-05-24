package main

import (
	"golang-ipc/pkg/ipc"
	"golang-ipc/pkg/types"
	"log"
	"net/rpc"
	"os"
)

type Client struct {
	conn *rpc.Client
}

func (c *Client) Dial() {
	c.conn = ipc.Dial()
}

func (c *Client) Call(method string, args interface{}, result interface{}) error {
	return ipc.Call(c.conn, method, args, result)
}

func main() {

	os.Setenv("IPC_PORT", "4040")

	client := &Client{}
	client.Dial()

	args := &types.Args{Msg: "hello"}
	resp := &types.Resp{}

	if err := client.Call("Server.PrintOut", args, resp); err != nil {
		log.Fatal("Server.PrintOut failed", err)
	}

	log.Println(resp.Success)
}