package parser

import (
	"io/ioutil"
	"testing"

	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/model"
)

const urlUserProfilePage = "http://localhost:8080/mock/album.zhenai.com/u/108415017"

func TestParseProfile(t *testing.T) {
	// contents, err := fetcher.Fetch(urlUserProfilePage)

	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("contents:%s", contents)

	result := parseProfile(contents, "惠儿", urlUserProfilePage)
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	actual := result.Items[0]

	expected := engine.Item{
		Url:  urlUserProfilePage,
		Type: "Zhenai",
		Id:   "108415017",
		Payload: model.Profile{
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
		}}
	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}

}
