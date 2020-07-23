package main

import (
	"github.com/go-crawler/zhenaiwang/engine"
	"github.com/go-crawler/zhenaiwang/persist"
	"github.com/go-crawler/zhenaiwang/rpcClient"
	"github.com/go-crawler/zhenaiwang/scheduler"
	"github.com/go-crawler/zhenaiwang/zhenai/parser"
)

func main() {
	saveClient := rpcClient.NewSaveClient(":12345")
	workerClient := rpcClient.NewWorkerClient(":123456")

	var requests []engine.Request
	requests = append(requests, engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSave(saveClient),
	}

	e.Run(requests...)
}
