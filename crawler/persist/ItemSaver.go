package persist

import (
	"errors"
	"log"

	"context"

	"github.com/ghjan/learngo/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {
	//Must turn off sniff in docker
	client, err := elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
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

			err := Save(client, item, index)
			if err != nil {
				log.Print("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) (error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id == "" {
		return errors.New("must supply Id")
	}
	_, err := indexService.Id(item.Id).Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v", resp)
	return err
}
