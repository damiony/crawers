package parser

import (
	"regexp"
	"strconv"

	"github.com/go-crawler/zhenaiwang/engine"
	"github.com/go-crawler/zhenaiwang/model"
)

func extraString(content []byte, matchRe string) string {
	re := regexp.MustCompile(matchRe)
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func ParseProfile(content []byte, name string) engine.ParseResult {
	var ageRe string = `<div class="m-btn purple" data-v-8b1eac0c>(\d+)岁</div>`
	var addrRe string = `<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`
	var marriageRe string = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
	var heightRe string = `<div class="m-btn purple" data-v-8b1eac0c>(\d+)cm</div>`
	var profile model.Profile

	profile.Name = name
	profile.Addr = extraString(content, addrRe)
	age, err := strconv.Atoi(extraString(content, ageRe))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extraString(content, marriageRe)
	height, err := strconv.Atoi(extraString(content, heightRe))
	if err == nil {
		profile.Height = height
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
