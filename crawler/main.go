package main

import (
	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/fetcher"
	"github.com/ghjan/learngo/crawler/persist"
	"github.com/ghjan/learngo/crawler/scheduler"
	"github.com/ghjan/learngo/crawler/zhenai/parser"
)

const (
	urlCityListPage = "http://www.zhenai.com/zhenghun"
	cityListRe      = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	//	cityText = `<a href="http://www.zhenai.com/zhenghun/aba" class="">阿坝</a>
	//<a href="http://www.zhenai.com/zhenghun/akesu" class="">阿克苏</a>
	//`
	urlShanghaiPage = "http://www.zhenai.com/zhenghun/shanghai"
)

func main() {
	//printCityList([]byte(cityText))
	//testCityList()
	//testSimpleEngine()
	testConcurentEngine()
	//testShanghai()
}

func testShanghai() {
	eng := engine.ConcurentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100}
	eng.Run(engine.Request{Url: urlShanghaiPage, Parser: engine.NewFuncParser(parser.ParseCity, " ParseCity")})
}
func testConcurentEngine() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	eng := engine.ConcurentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100, ItemChan: itemChan}
	eng.Run(engine.Request{Url: urlCityListPage, Parser: engine.NewFuncParser(parser.ParseCityList, " ParseCityList")})
}

func testSimpleEngine() {
	eng := engine.SimpleEngine{}
	eng.Run(engine.Request{Url: urlCityListPage, Parser: engine.NewFuncParser(parser.ParseCityList, " ParseCityList")})
}

func testCityList() {
	content, err := fetcher.Fetch(urlCityListPage)
	if err != nil {
		panic(err)
	}
	parser.ParseCityList(content, urlCityListPage)
}
