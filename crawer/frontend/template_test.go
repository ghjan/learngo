package frontend

import (
	"html/template"
	"os"
	"testing"

	"github.com/ghjan/learngo/crawer/engine"
	"github.com/ghjan/learngo/crawer/frontend/model"
	common "github.com/ghjan/learngo/crawer/model"
)

const urlUserProfilePage = "http://album.zhenai.com/u/108415017"

func TestTemplate(t *testing.T) {
	template := template.Must(template.ParseFiles("template.html"))
	page := model.SearchResult{}
	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	page.Hits = 123
	item := engine.Item{
		Url:  urlUserProfilePage,
		Type: "Zhenai",
		Id:   "108415017",
		Payload: common.Profile{
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
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}

}
