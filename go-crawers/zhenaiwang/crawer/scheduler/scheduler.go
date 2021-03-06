package scheduler

import "crawer/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.WorkerChan <- r
}

func (s *SimpleScheduler) ConfigurationWorkerChan(in chan engine.Request) {
	s.WorkerChan = in
}
