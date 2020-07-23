package model

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"task/parser"
)

type Profile struct {
	MemberId      int    // ID
	NickName      string // 姓名
	Sex           int    // 性别 0男 1女
	Height        int    // 身高
	Age           int    // 年龄
	Constellation string // 星座
	Education     string // 教育
	Marriage      string // 婚姻
}

type Item struct {
	Url    string
	Type   string
	Id     int
	Upload Profile
}

type RequestResult struct {
	Requests []Request
	Items    []Item
}

type Parser interface {
	Parse(contents []byte) RequestResult
	Serialize() (string, interface{})
}

type FuncParser struct {
	Name string
	Args interface{}
}

func (f *FuncParser) Parse(contents []byte) RequestResult {
	switch f.Name {
	case "ParseCityList":
		return parser.ParseCityList(contents)
	case "ParseCity":
		return parser.ParseCity(contents)
	default:
		fmt.Println("Invalid name:", f.Name)
		return RequestResult{}
	}
}

func (f *FuncParser) Serialize() (string, interface{}) {
	return f.Name, nil
}

type Request struct {
	Url    string
	Parser FuncParser
}

func ParseCityList(contents []byte) ParseResult {
	const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]+>([^<]+)</a>`
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)

	result := ParseResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}

func ParseCity(contents []byte) ParseResult {
	const userRe = `,"memberList":(\[[^\]]+"}\])`
	re := regexp.MustCompile(userRe)
	matchs := re.FindAllSubmatch(contents, -1)
	if matchs == nil {
		return ParseResult{}
	}

	var result ParseResult
	for _, m := range matchs {
		users := []Profile{}
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
