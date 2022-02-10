package parser

import (
	"fmt"
	"regexp"

	"github.com/ghjan/learngo/crawler/config"
	"github.com/ghjan/learngo/crawler/engine"
)

const cityListRe = `<a href="(http://.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, url string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllStringSubmatch(string(contents), -1)
	results := engine.ParseResult{}
	//limitCities := 10
	for _, m := range matches {
		//results.Items = append(results.Items, "City "+string(m[2]))
		results.Requests = append(results.Requests, engine.Request{Url: string(m[1]), Parser: engine.NewFuncParser(ParseCity, config.ParseCity)})
		fmt.Printf("-City:%s, URL:%s\n", m[2], m[1])
		//limitCities--
		//if limitCities <= 0 {
		//	break
		//}
	}
	fmt.Printf("ParseCityList, Matches found: %d\n", len(matches))
	return results
}
