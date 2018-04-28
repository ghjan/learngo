package main

import (
	"github.com/ghjan/learngo/crawler_distributed/config"
	"github.com/ghjan/learngo/crawler_distributed/persist"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"fmt"
)

func main() {
	serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ESIndex)
}
func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{Client: client, Index: index})
}
