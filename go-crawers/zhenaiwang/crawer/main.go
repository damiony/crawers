package main

import (
	"crawer/engine"
	"crawer/persist"
	"crawer/rpcClient"
	"crawer/scheduler"
	"crawer/worker"
	"crawer/zhenai/parser"
)

func main() {
	saveClient := rpcClient.NewSaveClient(":30000")
	worker := worker.CreateWorker(":30001")

	var requests []engine.Request
	requests = append(requests, engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.CreateFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSave(saveClient),
		Worker:      worker,
	}

	e.Run(requests...)
}
