package main

import (
	"errors"
	"fmt"

	"flag"
	"net/rpc"
	"strings"

	config1 "github.com/ghjan/learngo/crawler/config"
	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/fetcher"
	"github.com/ghjan/learngo/crawler/scheduler"
	"github.com/ghjan/learngo/crawler/zhenai/parser"
	itemsaverClient "github.com/ghjan/learngo/crawler_distributed/client"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
	workerClient "github.com/ghjan/learngo/crawler_distributed/worker/client"
)

const (
	urlCityListPage = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
	cityListRe      = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	//	cityText = `<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/aba" class="">阿坝</a>
	//<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu" class="">阿克苏</a>
	//`
	urlShanghaiPage = "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai"
)

var (
	itemSaverHost = flag.String("itemSaverHost", "", "item saver port")
	workerHosts   = flag.String("workerHosts", "", "crawler worker hosts seperated by comma")
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
	flag.Parse()
	if *itemSaverHost == "" {
		fmt.Println("please give a itemSaverHost!")
		return
	}
	if *workerHosts == "" {
		fmt.Println("please give a workerHosts!")
		return
	}
	itemChan, err := itemsaverClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool, err := createClientPool(strings.Split(*workerHosts, ","))
	if err != nil {
		panic(err)
	}
	processor := workerClient.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	eng := engine.ConcurentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
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

func createClientPool(hosts []string) (chan *rpc.Client, error) {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			fmt.Printf("Connected to %s\n", h)
		} else {
			fmt.Printf("Error connecting to %s\n", h)
			continue
		}
	}
	if len(clients) == 0 {
		return nil, errors.New(
			"no connections available")
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, nil
}
