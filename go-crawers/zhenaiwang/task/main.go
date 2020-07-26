package main

import "task/rpcServer"

func main() {
	rpcServer.RegisterWorkerService(":30001")
}
