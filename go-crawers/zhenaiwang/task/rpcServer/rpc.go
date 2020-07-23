package rpcServer

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"task/model"
)


func (s *SerializeParser) Worker)

func (s *SerializeParser) RegisterWorkerService(host string) {
	rpc.Register("WorkerService", new(SerializeParser))
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
