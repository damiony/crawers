package engine

import (
	"fmt"
	"log"

	"github.com/go-crawler/zhenaiwang/fetcher"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := s.worker(r)
		if err != nil {
			log.Println(err)
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			fmt.Printf("Got : %v\n", item)
		}
	}
}

func (s SimpleEngine) worker(r Request) (ParseResult, error) {
	fmt.Println("Url: ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)
	return parseResult, nil
}
