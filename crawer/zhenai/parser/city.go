package parser

import (
	"fmt"
	"regexp"

	"github.com/ghjan/learngo/crawer/engine"
)

var profileRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)">[^<]+</a>`)

func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllStringSubmatch(string(contents), -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		name := string(m[2])
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParseFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}

	matches2 := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches2 {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}

	fmt.Printf("ParseCity, Matches found: %d, matches2 found:%d \n", len(matches), len(matches2))
	return result
}
