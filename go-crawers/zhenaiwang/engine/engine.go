package engine

import (
	"fmt"
	"log"

	"github.com/go-crawler/zhenaiwang/fetcher"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		fmt.Println("Url: ", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Println(err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			fmt.Printf("Got : %v\n", item)
		}
	}
}
