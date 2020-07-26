package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
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

type Parser interface {
	Parse(contents []byte) RequestResult
	Serialize() (string, interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type RequestResult struct {
	Requests []Request
	Items    []Item
}

type ParserFunc func([]byte) RequestResult

type FuncParser struct {
	method ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte) RequestResult {
	return f.method(contents)
}

func (f *FuncParser) Serialize() (string, interface{}) {
	return f.name, nil
}

func CreateFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		method: p,
		name:   name,
	}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte) RequestResult {
	return RequestResult{}
}

func (NilParser) Serialize() (string, interface{}) {
	return "", nil
}

func ParseCityList(contents []byte) RequestResult {
	const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]+>([^<]+)</a>`
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)

	result := RequestResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, Request{
			Url:    string(m[1]),
			Parser: CreateFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}

func ParseCity(contents []byte) RequestResult {
	const userRe = `,"memberList":(\[[^\]]+"}\])`
	re := regexp.MustCompile(userRe)
	matchs := re.FindAllSubmatch(contents, -1)
	if matchs == nil {
		return RequestResult{}
	}

	var result RequestResult
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
				Item{
					Url:    fmt.Sprintf("https://album.zhenai.com/u/%d", user.MemberId),
					Type:   "zhenai",
					Id:     user.MemberId,
					Upload: user,
				})
		}
	}
	return result
}
