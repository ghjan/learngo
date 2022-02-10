package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/ghjan/learngo/crawler/config"
	rpcnames "github.com/ghjan/learngo/crawler_distributed/config"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
	"github.com/ghjan/learngo/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{}, nil)
	time.Sleep(2 * time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Use a fake fetcher to handle the url.
	// So we don't get data from zhenai.com
	const urlUserProfilePage = "http://localhost:8080/mock/album.zhenai.com/u/108415017"
	name := "惠儿"
	req := worker.Request{
		Url: urlUserProfilePage,
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: name,
		},
	}
	var result worker.ParseResult
	err = client.Call(
		rpcnames.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

	// Verify results
	assert.Equal(t, urlUserProfilePage, result.Items[0].Url)
	assert.Equal(t, name, result.Items[0].Payload.(map[string]interface{})["Name"])
}
