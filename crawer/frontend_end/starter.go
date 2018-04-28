package main

import (
	"net/http"

	"github.com/ghjan/learngo/crawer/frontend/controller"
	"fmt"
	"os"
)

var (
	templateFiles = []string{
		"crawer/frontend/view/template.html",
		"view/template.html",
	}
)

func main() {
	for _, filename := range templateFiles {
		if PathExist(filename) {
			fmt.Println(filename)
			http.Handle("/search", controller.CreateSearchResultHandler(filename))
		}
	}
	http.Handle("/", http.FileServer(
		http.Dir("crawer/frontend/view")))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
