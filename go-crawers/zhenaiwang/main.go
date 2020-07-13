package main

import (
	"fmt"

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
		Scheduler: &scheduler.SimpleScheduler{
			WorkerChan: make(chan engine.Request),
		},
		WorkerCount: 10,
	}
	fmt.Println(e.Scheduler)
	e.Run(requests...)
}
