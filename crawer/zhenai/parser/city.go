package parser

import (
	"github.com/ghjan/learngo/crawer/engine"
	"regexp"
	"fmt"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllStringSubmatch(string(contents), -1)
	results := engine.ParseResult{}
	for _, m := range matches {
		results.Items = append(results.Items, "User "+string(m[2]))
		results.Requests = append(results.Requests,
			engine.Request{
				Url: string(m[1]),
				ParseFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, string(m[2]))
				},
				})
	}
	fmt.Printf("ParseCity, Matches found: %d\n", len(matches))
	return results
}
