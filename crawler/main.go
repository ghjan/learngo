package main

import (
	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/fetcher"
	"github.com/ghjan/learngo/crawler/persist"
	"github.com/ghjan/learngo/crawler/scheduler"
	"github.com/ghjan/learngo/crawler/zhenai/parser"
	"github.com/ghjan/learngo/crawler/config"
)

const (
	urlCityListPage = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
	cityListRe      = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	//	cityText = `<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/aba" class="">阿坝</a>
	//<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu" class="">阿克苏</a>
	//`
	urlShanghaiPage = "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai"
)

func main() {
	//printCityList([]byte(cityText))
	//testCityList()
	//testSimpleEngine()
	testConcurentEngine()
	//testShanghai()
}

func testShanghai() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	eng := engine.ConcurentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100, ItemChan: itemChan,
		RequestProcessor: engine.Worker}
	eng.Run(engine.Request{Url: urlShanghaiPage, Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity)})
}
func testConcurentEngine() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	eng := engine.ConcurentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100, ItemChan: itemChan,
		RequestProcessor: engine.Worker}
	eng.Run(engine.Request{Url: urlCityListPage, Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)})
}

func testSimpleEngine() {
	eng := engine.SimpleEngine{}
	eng.Run(engine.Request{Url: urlCityListPage, Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)})
}

func testCityList() {
	content, err := fetcher.Fetch(urlCityListPage)
	if err != nil {
		panic(err)
	}
	parser.ParseCityList(content, urlCityListPage)
}
