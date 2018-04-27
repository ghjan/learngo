package persist

import (
	"testing"

	"github.com/ghjan/learngo/crawer/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
)

func TestItemSaver(t *testing.T) {
	expected := model.Profile{
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
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	// TODO :Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://elastic.davidzhang.xin:9200", "http://localhost:9200"),
		elastic.SetMaxRetries(10), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("Got %v, expected %v", actual, expected)
	}

}
