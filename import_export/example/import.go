package example

import (
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
)

//type people struct {
//	Firstname   string `json:"firstname"`
//	Lastname    string `json:"lastname"`
//	Institution string `json:"institution"`
//	Email       string `json:"email"`
//}
//
//type item struct {
//	Id       string   `json:"id"`
//	Title    string   `json:"title"`
//	Journal  string   `json:"journal"`
//	Volume   int      `json:"volume"`
//	Number   int      `json:"number"`
//	Pages    string   `json:"pages"`
//	Year     int      `json:"year"`
//	Authors  []people `json:"authors"`
//	Abstract string   `json:"abstract"`
//	Link     string   `json:"link"`
//	Keywords []string `json:"keywords"`
//	Body     string   `json:"body"`
//}

var client *elastic.Client

func initClient() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
}
func SaveItem(client *elastic.Client, item Course, index string, typ string) (err error) {
	var (
		indexService *elastic.IndexService
	)
	if typ != "" {
		indexService = client.Index().Index(index).Type(typ).BodyJson(item)
	} else {
		indexService = client.Index().Index(index).BodyJson(item)
	}
	if item.Id == "" {
		_, err = indexService.Id(item.Id).Do(context.Background())
	} else {
		_, err = indexService.Do(context.Background())
	}
	return
}

func BulkIndex(index string, typ string, data []Course) (successItems []Course, err error) {
	successItems = make([]Course, 0)
	if client == nil {
		initClient()
	}
	for _, item := range data {
		if err = SaveItem(client, item, index, typ); err != nil {
			return
		} else {
			successItems = append(successItems, item)
		}
	}
	return
}
