package main

import (
	"github.com/go-crawler/zhenaiwang/engine"
	"github.com/go-crawler/zhenaiwang/zhenai/parser"
)

func main() {
	var requests []engine.Request
	requests = append(requests, engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	engine.Run(requests...)
}
