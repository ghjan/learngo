package main

import (
	"fmt"
	"github.com/ghjan/learngo/lang/retriever/mock"
	"github.com/ghjan/learngo/lang/retriever/real"
	"time"

)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
	fmt.Stringer
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func Post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"contents": "another faked immooc.com",
	})
	return rp.Get(url)
}

func main() {
	fmt.Println("---mock.Retriever---------")
	var r Retriever
	mockerRetriever := mock.Retriever{"this is a fake imooc.com"}
	r = &mockerRetriever
	inspect(r)

	fmt.Println("---real.Retriever---------")
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// Type Assertion
	fmt.Println("-----Type Assertion-------")
	if rr, ok := r.(*mock.Retriever); ok {
		fmt.Println(rr.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}
	fmt.Println("-----session RetrieverPoster-------")
	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(&mockerRetriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting ", r)
	fmt.Printf(">%T %v\n", r, r)
	fmt.Print("Type Switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	default:
		fmt.Println("unknown Retriever type")
	}
}
