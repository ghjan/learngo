package main

import (
	"fmt"
	"github.com/ghjan/learngo/lang/retriever/interface_usage/infra"
	retriever2 "github.com/ghjan/learngo/lang/retriever/interface_usage/retriever"
	"github.com/ghjan/learngo/lang/retriever/interface_usage/testing"
)

func main() {
	var ()
	retrievers := make([]retriever2.Retriever, 0)
	retrievers = append(retrievers, infra.NewInfraRetriever())
	retrievers = append(retrievers, testing.NewTestingRetriever())
	for _, retriever := range retrievers {
		if result, err := retriever.Get("http://www.imooc.com"); err == nil {
			fmt.Println(result)
		} else {
			fmt.Errorf("error:%s", err.Error())
		}
	}
}
