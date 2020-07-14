package main

import (
	"github.com/go-crawler/zhenaiwang/engine"
	"github.com/go-crawler/zhenaiwang/scheduler"
	"github.com/go-crawler/zhenaiwang/zhenai/parser"
)

func main() {
	var requests []engine.Request
	requests = append(requests, engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(requests...)
}
