package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
	Worker      Worker
}

type Scheduler interface {
	Submit(Request)
	WorkderReady(chan Request)
	ConfigurationWorkerChan(chan Request)
	Run()
	CreateWokerChan() chan Request
}

func (c *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan RequestResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		go func() {
			in := make(chan Request)
			c.Worker(in, out, c.Scheduler.WorkderReady)
		}()
	}

	for _, r := range seed {
		c.Scheduler.Submit(r)
	}

	resultCounts := 0
	for {
		result := <-out
		resultCounts++
		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
		for _, item := range result.Items {
			c.ItemChan <- item
		}
	}
}
