package engine

import "github.com/ghjan/learngo/crawler/fetcher"

func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := r.Parser.Parse(body, r.Url)
	return parseResult, nil
}
