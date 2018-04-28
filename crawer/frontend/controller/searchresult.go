package controller

import (
	"net/http"
	"strings"

	"github.com/ghjan/learngo/crawer/frontend/view"
	"github.com/olivere/elastic"
	"strconv"
	"fmt"
	"github.com/ghjan/learngo/crawer/frontend/model"
	"context"
	"reflect"
	"github.com/ghjan/learngo/crawer/engine"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//ServeHTTP localhost:8888/search?q=男 已购房&from=20  这里的from是开始的记录（因为分页）
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	fmt.Printf("q:%s\n", q)
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Printf("from:%d\n", from)
	fmt.Fprintf(w, "q=%s, from=%d", q, from)
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("dating_profile").Type("zhenai").Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	for _, item := range resp.Each(reflect.TypeOf(engine.Item{})) {
		result.Items = append(result.Items, item.(engine.Item))
	}

	return result, err

}
