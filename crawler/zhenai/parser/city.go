package parser

import (
	"fmt"
	"regexp"

	"github.com/ghjan/learngo/crawler/config"
	"github.com/ghjan/learngo/crawler/engine"
)

var profileRe = regexp.MustCompile(`<a href="(http://.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://.*www\.zhenai\.com/zhenghun/[^"]+)">[^<]+</a>`)

func ParseCity(contents []byte, url string) engine.ParseResult {

	matches := profileRe.FindAllStringSubmatch(string(contents), -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "User "+string(m[2]))
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    url,
				Parser: NewProfileParser(name),
			})
		//fmt.Printf("--ParseCity, profile url:%s \n", url)
	}

	matches2 := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches2 {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
		//fmt.Printf("--ParseCity, city url:%s \n", string(m[1]))
	}

	fmt.Printf("ParseCity, Matches(profile) found: %d, matches2(city) found:%d \n", len(matches), len(matches2))
	return result
}
