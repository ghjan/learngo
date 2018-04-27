package persist

import (
	"errors"
	"log"

	"context"

	"github.com/ghjan/learngo/crawer/engine"
	"github.com/olivere/elastic"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Print("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) (error) {
	//Must turn off sniff in docker
	client, err := elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)
	if item.Id == "" {
		return errors.New("must supply Id")
	}
	_, err = indexService.Id(item.Id).Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v", resp)
	return err
}
