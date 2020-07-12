package engine

import (
	"fmt"
	"log"

	fetch "github.com/go-crawler/zhenaiwang/fetcher"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	fmt.Println("requests len is:", len(requests))
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetch.Fetch(r.Url)
		if err != nil {
			log.Println(err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			fmt.Printf("Got item: %v", item)
		}
	}
}
