package main

import (
	"fmt"
	"testing"

	"time"

	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/model"
	"github.com/ghjan/learngo/crawler_distributed/config"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
)

const urlUserProfilePage = "http://localhost:8080/mock/album.zhenai.com/u/108415017"
const index = "test1"
const ItemSaverServerDefaultPort = 1234

func TestItemSaver(t *testing.T) {
	serviceHost := fmt.Sprintf(":%d", ItemSaverServerDefaultPort)
	// start ItemSaverServer
	var (
		serverReady = make(chan struct{})
		err         error
	)
	go func() {
		if err = serveRpc(serviceHost, index, serverReady); err != nil {
			t.Errorf("failed, serveRpc error:%s", err.Error())
		}
	}()
	<-serverReady
	time.Sleep(time.Second)
	// start ItemSaverClient

	client, err := rpcsupport.NewClient(serviceHost)
	if err != nil {
		t.Fatalf("failed, error:%s", err.Error())
	}

	// Call save
	item := engine.Item{
		Url:  urlUserProfilePage,
		Type: "Zhenai",
		Id:   "108415017",
		Payload: model.Profile{
			Age:        50,
			Height:     156,
			Weight:     0,
			Income:     "3000元以下",
			Gender:     "女",
			Name:       "惠儿",
			Xinzuo:     "魔羯座",
			Occupation: "销售总监",
			Marriage:   "离异",
			House:      "租房",
			Hokou:      "四川阿坝",
			Education:  "高中及以下",
			Car:        "未购车",
		}}

	result := ""
	err = client.Call(config.ItemSaverRpc,
		item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s",
			result, err)
		t.Fail()
	}
}
