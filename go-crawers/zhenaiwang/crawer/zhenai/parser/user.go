package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/go-crawler/zhenaiwang/engine"
	"github.com/go-crawler/zhenaiwang/model"
)

const userRe = `,"memberList":(\[[^\]]+"}\])`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userRe)
	matchs := re.FindAllSubmatch(contents, -1)
	if matchs == nil {
		return engine.ParseResult{}
	}

	var result engine.ParseResult
	for _, m := range matchs {
		users := []model.Profile{}
		err := json.Unmarshal(m[1], &users)
		if err != nil {
			log.Println("json unmarshal error:", err)
			fmt.Println("input: ", string(m[1]))
			continue
		}
		for _, user := range users {
			result.Items = append(result.Items,
				engine.Item{
					Url:    fmt.Sprintf("https://album.zhenai.com/u/%d", user.MemberId),
					Type:   "zhenai",
					Id:     user.MemberId,
					Upload: user,
				})
		}
	}
	return result
}
