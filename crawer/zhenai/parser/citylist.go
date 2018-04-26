package parser

import (
	"github.com/ghjan/learngo/crawer/engine"
	"regexp"
	"fmt"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllStringSubmatch(string(contents), -1)
	results := engine.ParseResult{}
	for _, m := range matches {
		results.Items = append(results.Items, string(m[2]))
		results.Requests = append(results.Requests, engine.Request{Url: string(m[1]), ParseFunc: engine.NilParser})
		//fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
	}
	fmt.Printf("ParseCityList, Matches found: %d\n", len(matches))
	return results
}
