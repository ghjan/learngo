package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ghjan/learngo/crawler_distributed/config"
	"github.com/ghjan/learngo/crawler_distributed/persist"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "itemsaver port")

func main() {
	var (
		err error
	)
	flag.Parse()
	if *port == 0 {
		fmt.Println("please give a port!")
		return
	}
	fmt.Println(*port)
	if err = serveRpc(fmt.Sprintf(":%d", *port), config.ESIndex, nil); err != nil {
		log.Fatal(err)
	}
}

func serveRpc(host string, index string, serverReady chan struct{}) (err error) {
	var (
		client *elastic.Client
	)
	client, err = elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{Client: client, Index: index}, serverReady)
}
