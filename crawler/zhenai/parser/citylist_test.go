package parser

import (
	"testing"

	"io/ioutil"
	//"github.com/ghjan/learngo/crawler/engine"
)

const urlCityListPage = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func TestParseCityList(t *testing.T) {

	//contents, err := fetcher.Fetch(urlCityListPage)
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("contents:%s", contents)
	result := ParseCityList(contents, urlCityListPage)
	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	//expectedCities := []engine.Item{
	//	{Id: "aba", Type: "zhenai", Payload: "City 阿坝"},
	//	{Id: "akesu", Type: "zhenai", Payload: "City 阿克苏"},
	//	{Id: "alashanmeng", Type: "zhenai", Payload: "City 阿拉善盟"},
	//}

	for len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d", resultSize, len(result.Requests))
	}
	//if len(result.Items) != resultSize {
	//	t.Errorf("result should have %d "+"items; but had %d", resultSize, len(result.Items))
	//}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected city #%d: %s; but "+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}

	//for i, city := range expectedCities {
	//	if result.Items[i] != city {
	//		t.Errorf("expected city #%d: %s; but "+
	//			"was %s",
	//			i, city, result.Items[i])
	//	}
	//}

}
