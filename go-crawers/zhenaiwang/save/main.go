package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"save/rpcServer"
)

func main() {
	rpcServer.RegisterItemSaveService()
	listen, err := net.Listen("tcp", ":30000")
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
