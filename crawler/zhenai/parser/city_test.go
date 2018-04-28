package parser

import (
	"testing"

	//"github.com/ghjan/learngo/crawler/fetcher"
	"io/ioutil"
	//"github.com/ghjan/learngo/crawler/engine"
)

const urlCityUserPage = "http://www.zhenai.com/zhenghun/aba"

func TestParseCity(t *testing.T) {

	//contents, err := fetcher.Fetch(urlCityUserPage)
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	//contents = []byte(userContent)
	//fmt.Printf("contents:%s", contents)
	result := ParseCity(contents, urlCityUserPage)
	const resultSize = 20
	expectedUrls := []string{
		"http://album.zhenai.com/u/108415017",
		"http://album.zhenai.com/u/1314495053",
		"http://album.zhenai.com/u/1121586032",
	}
	//expectedUsers := []engine.Item{
	//	{Id: "108415017", Type: "Zhenai", Payload: "User 惠儿"},
	//	{Id: "1314495053", Type: "Zhenai", Payload: "User 风中的蒲公英"},
	//	{Id: "1121586032", Type: "Zhenai", Payload: "User 现实与理想之间"},
	//}

	for len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d", resultSize, len(result.Requests))
	}
	//if len(result.Items) != resultSize {
	//	t.Errorf("result should have %d "+"items; but had %d", resultSize, len(result.Items))
	//}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}

	//for i, user := range expectedUsers {
	//	if result.Items[i] != user {
	//		t.Errorf("expected user #%d: %s; but "+
	//			"was %s",
	//			i, user, result.Items[i])
	//	}
	//}

}
