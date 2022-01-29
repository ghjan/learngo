package testing

import "github.com/ghjan/learngo/lang/retriever/interface_usage/retriever"

type TestingRetriever struct {
}

func (TestingRetriever) Get(url string) (result string, err error) {
	result = "fake result"
	return
}

func NewTestingRetriever() retriever.Retriever {
	return TestingRetriever{}
}

