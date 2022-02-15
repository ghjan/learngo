package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"os"
	"strings"

	"path/filepath"

	"github.com/ghjan/learngo/crawler/frontend/controller"
)

var (
	templateFiles = []string{
		"crawler/frontend/view/template.html",
		"view/template.html",
	}
)
var port = flag.Int("port", 0, "frontend server port")

func main() {
	var (
		err error
	)
	flag.Parse()
	if *port == 0 {
		fmt.Println("please give a port!")
		return
	}

	pathPrefix := "crawler/front/view"
	for _, filename := range templateFiles {
		if PathExist(filename) {
			pathPrefix = getPath(filename)
			http.Handle("/search", controller.CreateSearchResultHandler(filename))
			break
		}
	}

	http.Handle("/", http.FileServer(http.Dir(pathPrefix)))
	host := fmt.Sprintf(":%d", *port)
	fmt.Printf("frontend server listening on %s\n", host)
	err = http.ListenAndServe(host, nil)

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
