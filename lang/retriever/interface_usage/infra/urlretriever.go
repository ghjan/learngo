package infra

import (
	"github.com/ghjan/learngo/lang/retriever/interface_usage/retriever"
	"io/ioutil"
	"net/http"
)

type InfraRetriever struct {
}

func (InfraRetriever) Get(url string) (result string, err error) {
	var (
		resp *http.Response
	)
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	result = string(bytes)
	return
}

func NewInfraRetriever() retriever.Retriever {
	return InfraRetriever{}
}

