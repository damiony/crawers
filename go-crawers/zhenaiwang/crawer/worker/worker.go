package worker

import (
	"crawer/engine"
	"crawer/rpcClient"
	"crawer/zhenai/parser"
	"fmt"
	"log"
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

func DeserializeParser(s SerializeParser) (engine.Parser, error) {
	switch s.Name {
	case "ParseCityList":
		return engine.CreateFuncParser(parser.ParseCityList, "ParseCityList"), nil
	case "ParseCity":
		return engine.CreateFuncParser(parser.ParseCity, "ParseCity"), nil
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
		//fmt.Println(err)
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

func CreateWorker(host string) engine.Worker {
	client := rpcClient.NewWorkerClient(host)
	return func(in chan engine.Request, out chan engine.RequestResult, ready func(chan engine.Request)) {
		for {
			ready(in)
			request := <-in
			//fmt.Println(request.Url)
			sRequest := SerializeRequest(request)
			var sRequestResult RequestResult
			err := client.Call("WorkerService.Worker", sRequest, &sRequestResult)
			if err != nil {
				log.Println(err)
				continue
			}

			requestResult := DeserializeRequestResult(sRequestResult)
			out <- requestResult
			//fmt.Println("ok:", request.Url)
		}
	}
}
