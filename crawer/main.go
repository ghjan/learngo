package main

import (
	"github.com/ghjan/learngo/crawer/fetcher"
	"github.com/ghjan/learngo/crawer/zhenai/parser"
	"github.com/ghjan/learngo/crawer/engine"
	"github.com/ghjan/learngo/crawer/scheduler"
)

const (
	urlCityListPage = "http://www.zhenai.com/zhenghun"
	cityListRe      = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	//	cityText = `<a href="http://www.zhenai.com/zhenghun/aba" class="">阿坝</a>
	//<a href="http://www.zhenai.com/zhenghun/akesu" class="">阿克苏</a>
	//`
)

func main() {
	//printCityList([]byte(cityText))
	//testCityList()
	//testSimpleEngine()
	testConcurentEngine()
}

func testConcurentEngine() {
	eng := engine.ConcurentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkerCount: 100}
	eng.Run(engine.Request{Url: urlCityListPage, ParseFunc: parser.ParseCityList})
}

func testSimpleEngine() {
	eng := engine.SimpleEngine{}
	eng.Run(engine.Request{Url: urlCityListPage, ParseFunc: parser.ParseCityList})
}

func testCityList() {
	content, err := fetcher.Fetch(urlCityListPage)
	if err != nil {
		panic(err)
	}
	parser.ParseCityList(content)
}
