package main

import (
	"log"
	"net/http"

	"fmt"
	"os"
	"strings"

	"path/filepath"

	"github.com/ghjan/learngo/crawer/frontend/controller"
)

var (
	templateFiles = []string{
		"crawer/frontend/view/template.html",
		"view/template.html",
	}
)

func main() {
	pathPrefix := "crawer/front/view"
	for _, filename := range templateFiles {
		if PathExist(filename) {
			pathPrefix = getPath(filename)
			fmt.Println(pathPrefix)
			http.Handle("/search", controller.CreateSearchResultHandler(filename))
			break
		}
	}

	http.Handle("/", http.FileServer(http.Dir(pathPrefix)))
	err := http.ListenAndServe(":8088", nil)

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

func getPath(filename string) string {
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
