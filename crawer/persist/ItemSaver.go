package persist

import (
	"log"

	"github.com/olivere/elastic"
	"context"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Print("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return out
}

func save(item interface{}) (string, error) {
	//Must turn off sniff in docker
	client, err := elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	//fmt.Printf("%+v", resp)
	return resp.Id, err
}
