package model

import "github.com/ghjan/learngo/crawer/engine"

type SearchResult struct {
	Hits  int64
	Start int
	Items []engine.Item
}
