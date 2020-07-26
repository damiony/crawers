package rpcServer

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"task/Worker"
)

func RegisterWorkerService(host string) {
	rpc.RegisterName("WorkerService", new(Worker.WorkerService))
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
