package engine

import "crawer/model"

type Parser interface {
	Parse(contents []byte) RequestResult
	Serialize() (string, interface{})
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

func CreateFuncParser(p ParserFunc, name string) Parser {
	return &FuncParser{
		method: p,
		name:   name,
	}
}

type Request struct {
	Url    string
	Parser Parser
}

type RequestResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url    string
	Type   string
	Id     int
	Upload model.Profile
}

func NilParse([]byte) RequestResult {
	return RequestResult{}
}

type Worker func(in chan Request, out chan RequestResult, ready func(chan Request))
