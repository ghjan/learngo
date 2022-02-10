package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}
type Scheduler interface {
	readyNotify
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type Processor func(Request) (ParseResult, error)

type readyNotify interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	//运行调度器
	e.Scheduler.Run()

	//创建若干个worker
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	//提交种子url（爬虫开始爬取信息的初始页面）
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		//处理爬取结果中的信息
		result := <-out
		for _, item := range result.Items {
			itemCount++
			fmt.Printf("Got item #%d:%v\n", itemCount, item)
			// TODO:item去重
			go func() {
				e.ItemChan <- item
			}()
		}

		//处理爬取结果中的其他url
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, notify readyNotify) {
	//in := make(chan Request)
	go func() {
		for {
			// tell notify i'm ready
			notify.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
