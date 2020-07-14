package scheduler

import (
	"github.com/go-crawler/zhenaiwang/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkderReady(work chan engine.Request) {
	q.workerChan <- work
}

func (q *QueuedScheduler) CreateWokerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) ConfigurationWorkerChan(in chan engine.Request) {
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var request engine.Request
			var workerChan chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				request = requestQ[0]
				workerChan = workerQ[0]
			}

			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case worker := <-q.workerChan:
				workerQ = append(workerQ, worker)
			case workerChan <- request:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}
	}()
}
