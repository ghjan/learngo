package persist

import (
	"log"

	"github.com/ghjan/learngo/crawler/engine"
	persistReal "github.com/ghjan/learngo/crawler/persist"
	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persistReal.Save(s.Client, item, s.Index)
	if err == nil {
		log.Printf("Item %v saved ", item)
		*result = "ok"
	} else {
		log.Printf("Error saving %v: %v", item, err)
	}
	return err
}
