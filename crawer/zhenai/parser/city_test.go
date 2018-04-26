package parser

import (
	"testing"

	"io/ioutil"
	//"fmt"
)

//const urlCityPage = "http://www.zhenai.com/zhenghun"
const userContent = `<li class="clearfix">
									<div class="guess-photo">
										<a href="http://album.zhenai.com/u/109816757" target="_blank">
											<img src="http://photo16.zastatic.com/images/photo/27455/109816757/99718035787672941.png?scrop=1&crop=1&w=100&h=100&cpos=north">
										</a>
									</div>
									<dl class="guess-info">
										<dt class="guess-name fs16">
											<a class="exp-user-name" target="_blank"
												href="http://album.zhenai.com/u/109816757">蝴蝶在飞舞</a>
										</dt>
										<dd>
											<p class="guess-age lh24 fs12 c9f">31岁 162cm</p>
											<p class="guess-gb fs14 c5e">我是东阳人，喜欢另一半也是东阳或义乌本地人。只要是我喜欢的哪里都好。</p>
										</dd>
									</dl>
								</li>`

func TestParseCity(t *testing.T) {

	//contents, err := fetcher.Fetch(urlCityPage)
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	//contents = []byte(userContent)
	//fmt.Printf("contents:%s", contents)
	result := ParseCity(contents)
	const resultSize = 2
	expectedUrls := []string{
		"http://album.zhenai.com/u/109816757",
		"http://album.zhenai.com/u/101873762",
	}
	expectedUsers := []string{
		"User 蝴蝶在飞舞", "User 燕子",
	}

	for len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"items; but had %d", resultSize, len(result.Items))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}

	for i, user := range expectedUsers {
		if result.Items[i].(string) != user {
			t.Errorf("expected user #%d: %s; but "+
				"was %s",
				i, user, result.Items[i])
		}
	}

}
