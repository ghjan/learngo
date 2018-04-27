package engine

import (
	"fmt"

	"github.com/ghjan/learngo/crawer/model"
)

type ConcurentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	readyNotify
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type readyNotify interface {
	WorkerReady(chan Request)
}

func (e *ConcurentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				profileCount++
				fmt.Printf("Got profile #%d:%v\n", profileCount, item)
			}
		}
		//URL dedup
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicat request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

func (e *ConcurentEngine) createWorker(in chan Request, out chan ParseResult, notify readyNotify) {
	//in := make(chan Request)
	go func() {
		for {
			// tell notify i'm ready
			notify.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
