package main

import (
	"fmt"
	"log"

	"github.com/ghjan/learngo/crawler_distributed/config"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
	"github.com/ghjan/learngo/crawler_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkPort0), worker.CrawlService{}))
}
