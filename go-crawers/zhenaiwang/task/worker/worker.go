package Worker

import "net/rpc"

type FuncParser struct {
	Name string
	Args interface{}
}

type SerializeParser struct {
	Url    string
	Parser model.Parser
}

type 

type WorkerService struct {}

func (w *WorkerService) Worker(s SerializeParser, reply *)

func RegisterWorkerService (host string) {
	rpc.Register("WorkerService", new(WorkerService))
}
