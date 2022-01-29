package infra

import "testing"

func TestRetriever_Get(t *testing.T) {
	retriever := Retriever{}
	if result, err := retriever.Get("http://www.imooc.com"); err == nil {
		t.Logf(result)
	} else {
		t.Errorf(err.Error())
	}
}
