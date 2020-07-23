package parser

import (
	"regexp"

	"github.com/go-crawler/zhenaiwang/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		// result.Items = append(result.Items, string(m[2]))
	}

	return result
}
