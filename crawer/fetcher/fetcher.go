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
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//提交请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	//增加header选项
	req.Header.Add("Cookie", "Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1524713461; sid=A1218s9FBBOVwXcCrafX; ipCityCode=10103000; ipOfflineCityCode=10103000; JSESSIONID=abczR5SeLN_i86H6M-amw; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1524737141")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	<-rateLimiter
	//resp, err := http.Get(url)
	resp, err := client.Do(req)
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
