package engine

import (
	"fmt"
	"log"

	"github.com/go-crawler/zhenaiwang/fetcher"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

func (c ConcurrentEngine) Run(seed ...Request) {
	for _, r := range seed {
		c.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)

	for i := 0; i < c.WorkerCount; i++ {
		go func() {
			c.createWorker(in, out)
		}()
	}

	for {
		parseResult := <-out
		for _, r := range parseResult.Requests {
			c.Scheduler.Submit(r)
		}
		for _, item := range parseResult.Requests {
			fmt.Printf("Got : %v\n", item)
		}
	}
}

func (c ConcurrentEngine) worker(r Request) (ParseResult, error) {
	fmt.Println("Url: ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)
	return parseResult, nil
}

func (c ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	for {
		r := <-in
		parseResult, err := c.worker(r)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- parseResult
	}
}
