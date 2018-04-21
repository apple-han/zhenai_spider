package scheduler

import (
	"learn/crawler/engine"

	"github.com/micro/go-micro/selector"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func(s QueuedScheduler) Submit(r engine.Request){
	s.requestChan <- r
}

func(s QueuedScheduler) WorkerReady(
	w chan engine.Request){
		s.workerChan <- w
}
func (s QueuedScheduler) ConfigureMasterWorkerChan(chan ){
	panic("implement me")
}

func(s QueuedScheduler) Run(){
	go func(){
		var requestQ []engine.Request
		var workerQ  []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if (len(requestQ) > 0 &&
				len(workerQ) > 0) {
				 activeWorker = workerQ[0]
				 activeRequest = requestQ[0]
			}
			select{
			case r := <- s.requestChan:
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}