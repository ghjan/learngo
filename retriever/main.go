package main

import (
	"fmt"
	"github.com/ghjan/learngo/retriever/mock"
	"github.com/ghjan/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	mockerRetriever := mock.Retriever{"this is a fake imooc.com"}
	r = &mockerRetriever
	inspect(r)

	fmt.Println("------------")
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// Type Assertion
	fmt.Println("------------")
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	//fmt.Println("Try a session with mockRetriever")
	//fmt.Println(session(&mockRetriever))

}
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type Switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	default:
		fmt.Println("unknown Retriever type")
	}
}
