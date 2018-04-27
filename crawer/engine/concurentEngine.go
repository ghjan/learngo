package engine

import (
	"fmt"
)

type ConcurentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
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
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			itemCount++
			fmt.Printf("Got item #%d:%v\n", itemCount, item)
			// TODO:item去重
			go func() {
				e.ItemChan <- item
			}()
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
