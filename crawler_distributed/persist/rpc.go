package persist

import (
	"github.com/ghjan/learngo/crawler/engine"
	persistReal "github.com/ghjan/learngo/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persistReal.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
