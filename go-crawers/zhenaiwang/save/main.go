package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"save/rpcServer"
	"time"
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
		time.Sleep(time.Second)
	}
}
