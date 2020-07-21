package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"

	"github.com/crawers/go-crawers/zhenaiwang/save/rpcServer"
)

func main() {
	rpcServer.RegisterItemSaveService()
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go jsonrpc.ServeConn(conn)
	}
}
