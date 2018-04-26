package fetcher

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyReader := bufio.NewReader(resp.Body)
		e := determingEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s\n", string(all))
		return all, err
	} else {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
}

func determingEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
