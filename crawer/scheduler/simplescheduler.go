package scheduler

import "github.com/ghjan/learngo/crawer/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send r down to worker chan
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

//func (s *SimpleScheduler) Run() {
//	s.workerChan = make(chan engine.Request)
//}