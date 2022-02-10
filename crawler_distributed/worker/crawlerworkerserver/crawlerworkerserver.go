package main

import (
	"fmt"
	"log"

	"flag"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
	"github.com/ghjan/learngo/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "crawler worker port")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("please give a port!")
		return
	}
	fmt.Println(*port)
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}, nil))
}
