package client

import (
	"log"

	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler_distributed/config"
	"github.com/ghjan/learngo/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			//Call RPC to save item
			result := ""
			err = client.Call(config.ItemSaverRpc,
				item, &result)

			if err != nil || result != "ok" {
				log.Printf("ItemSaver result: %s; err: %s",
					result, err)
			}
		}
	}()
	return out, nil
}
