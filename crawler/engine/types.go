package engine

import "github.com/ghjan/learngo/crawler/config"

type ParserFunc func([]byte, string) ParseResult

// Parser : 解析器
type Parser interface {
	// Parse : 解析方法 返回解析结果
	Parse(contents []byte, url string) ParseResult
	// Serialize : 序列化函数 返回序列化的名称，以及参数列表
	Serialize() (name string, args interface{})
}
type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// NewFuncParser :a factory to produce FuncParser
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
