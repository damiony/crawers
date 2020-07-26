package parser

import (
	"crawer/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) engine.RequestResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)

	result := engine.RequestResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//Parser: engine.CreateFuncParser(ParseCity, "ParseCity"),
		})
		// result.Items = append(result.Items, string(m[2]))
	}

	return result
}
