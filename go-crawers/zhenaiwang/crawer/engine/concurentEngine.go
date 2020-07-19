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
	WorkderReady(chan Request)
	ConfigurationWorkerChan(chan Request)
	Run()
	CreateWokerChan() chan Request
}

func (c *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		go func() {
			c.createWorker(out)
		}()
	}

	for _, r := range seed {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
		for _, item := range result.Items {
			fmt.Printf("Got : %v\n", item)
		}
	}
}

func (c *ConcurrentEngine) worker(r Request) (ParseResult, error) {
	fmt.Println("Url: ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)
	return parseResult, nil
}

func (c *ConcurrentEngine) createWorker(out chan ParseResult) {
	in := c.Scheduler.CreateWokerChan()
	for {
		c.Scheduler.WorkderReady(in)

		r := <-in
		result, err := c.worker(r)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- result
	}
}
