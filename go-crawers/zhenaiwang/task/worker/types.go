package Worker

import (
	"fmt"
	"task/engine"
	"task/fetcher"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializeParser
}

type RequestResult struct {
	Items    []engine.Item
	Requests []Request
}

type WorkerService struct{}

func (c *WorkerService) Worker(r Request, result *RequestResult) error {
	fmt.Printf("url: %s\n", r.Url)
	request := DeserilizeRequest(r)

	contents, err := fetcher.Fetch(request.Url)
	if err != nil {
		return err
	}

	requestResult := request.Parser.Parse(contents)

	*result = SerializeRequestResult(requestResult)
	return nil
}

func DeserializeParser(s SerializeParser) (engine.Parser, error) {
	switch s.Name {
	case "ParseCityList":
		return engine.CreateFuncParser(engine.ParseCityList, "ParseCityList"), nil
	case "ParseCity":
		return engine.CreateFuncParser(engine.ParseCity, "ParseCity"), nil
	case "Nil":
		return engine.NilParser{}, nil
	default:
		return nil, fmt.Errorf("Invalid name: %s", s.Name)
	}
}

func DeserilizeRequest(r Request) engine.Request {
	request := engine.Request{
		Url: r.Url,
	}

	parser, err := DeserializeParser(r.Parser)
	if err != nil {
		fmt.Println(err)
		return request
	}

	request.Parser = parser
	return request
}

func SerializeRequest(r engine.Request) Request {
	request := Request{
		Url: r.Url,
	}
	name, args := r.Parser.Serialize()
	request.Parser = SerializeParser{
		Name: name,
		Args: args,
	}

	return request
}

func DeserializeRequestResult(result RequestResult) engine.RequestResult {
	r := engine.RequestResult{
		Items: result.Items,
	}
	for _, request := range result.Requests {
		r.Requests = append(r.Requests, DeserilizeRequest(request))
	}
	return r
}

func SerializeRequestResult(result engine.RequestResult) RequestResult {
	r := RequestResult{
		Items: result.Items,
	}
	for _, request := range result.Requests {
		r.Requests = append(r.Requests, SerializeRequest(request))
	}
	return r
}
