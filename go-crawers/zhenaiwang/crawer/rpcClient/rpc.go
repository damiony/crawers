package rpcClient

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func NewSaveClient(host string) *rpc.Client {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	client := jsonrpc.NewClient(conn)

	return client
}

func NewWorkerClient(host string) *rpc.Client {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	client := jsonrpc.NewClient(conn)

	return client
}
