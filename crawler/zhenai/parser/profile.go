package parser

import (
	"regexp"
	"strconv"

	"github.com/ghjan/learngo/crawler/config"
	"github.com/ghjan/learngo/crawler/engine"
	"github.com/ghjan/learngo/crawler/model"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

//var guessRe = regexp.MustCompile(
//	`<a class="exp-user-name"[^>]*href="(http://*.album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`http://.*album\.zhenai\.com/u/([\d]+)`)

var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://.*album\.zhenai\.com/u/[0-9]+)">([^<]+)</a>`)

func parseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}

	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err == nil {
		profile.Weight = weight
	}

	profile.Name = name
	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	id := extractString([]byte(url), idUrlRe)

	result := engine.ParseResult{
		Items: []engine.Item{{Url: url, Type: "Zhenai", Id: id, Payload: profile}},
	}
	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    url,
				Parser: NewProfileParser(name),
			})

	}
	return result

}

//extractString 获取某个字符串
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, p.userName, url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

//func ProfileParser(name string) engine.ParserFunc {
//	return func(c []byte, url string) engine.ParseResult {
//		return parseProfile(c, name, url)
//	}
//}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{userName: name}
}
