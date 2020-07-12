package parser

import (
	"regexp"

	"github.com/go-crawler/zhenaiwang/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/\d+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	match := re.FindAllSubmatch(contents, -1)
	if match == nil {
		return engine.ParseResult{}
	}

	result := engine.ParseResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			},
		})
		result.Items = append(result.Items, "City "+string(m[2]))
	}

	return result
}
