package main

import (
	"fmt"

	config1 "github.com/ghjan/learngo/crawler/config"
	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/fetcher"
	"github.com/ghjan/learngo/crawler/scheduler"
	"github.com/ghjan/learngo/crawler/zhenai/parser"
	itemsaverClient "github.com/ghjan/learngo/crawler_distributed/client"
	workerClient "github.com/ghjan/learngo/crawler_distributed/worker/client"
	"github.com/ghjan/learngo/crawler_distributed/config"
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
	eng.Run(engine.Request{Url: urlShanghaiPage, Parser: engine.NewFuncParser(parser.ParseCity, config1.ParseCity)})

}
func testConcurentEngine() {
	itemChan, err := itemsaverClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := workerClient.CreateProcessor()
	if err != nil {
		panic(err)
	}
	eng := engine.ConcurentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100,
		ItemChan: itemChan, RequestProcessor: processor}
	eng.Run(engine.Request{Url: urlCityListPage,
		Parser: engine.NewFuncParser(parser.ParseCityList, config1.ParseCityList)})
}

func testSimpleEngine() {
	eng := engine.SimpleEngine{}
	eng.Run(engine.Request{Url: urlCityListPage,
		Parser: engine.NewFuncParser(parser.ParseCityList, config1.ParseCityList)})
}

func testCityList() {
	content, err := fetcher.Fetch(urlCityListPage)
	if err != nil {
		panic(err)
	}
	parser.ParseCityList(content, urlCityListPage)
}
