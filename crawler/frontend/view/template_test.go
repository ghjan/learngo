package view

import (
	"os"
	"testing"

	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/frontend/model"
	common "github.com/ghjan/learngo/crawler/model"
)

const urlUserProfilePage = "http://album.zhenai.com/u/108415017"

func TestTemplate(t *testing.T) {
	//template := template.Must(template.ParseFiles("template.html"))
	view := CreateSearchResultView("template.html")
	pageData := model.SearchResult{}
	pageData.Hits = 123
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
		pageData.Items = append(pageData.Items, item)
	}
	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}

	err = view.Render(out, pageData)

	if err != nil {
		panic(err)
	}

}
