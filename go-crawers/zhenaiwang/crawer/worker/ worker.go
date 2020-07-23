package worker

import (
	"fmt"
	"net/rpc"

	"github.com/go-crawler/zhenaiwang/engine"
)

func CreateWorker() chan chan engine.Request {
	in := make(chan chan engine.Request)
			for {
			r := 
			}
}

func Worker(r engine.Request, client *rpc.Client) {
	var reply string
	err := client.Call("WorkerService.Worker", r, reply)
	if err != nil {
		fmt.Printf("Rpc call worker error: %+v\n", err)
	}
}
