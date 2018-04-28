package main

import (
	"gopkg.in/olivere/elastic.v5"
	"github.com/ghjan/learngo/crawler_distributed/persist"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
)

func main() {
	serveRpc(":1234", "dating_profile")
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
