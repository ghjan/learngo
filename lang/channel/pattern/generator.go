package pattern

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func MsgGen(name string, tickSeconds int, done chan struct{}, cleanupConsuming int) chan string {
	c := make(chan string)
	if tickSeconds <= 0 { //默认为不超过两秒钟
		tickSeconds = 2
	}
	go func() {
		i := 0
		for {
			if done != nil {
				select {
				case <-time.After(time.Duration(rand.Intn(1000)*tickSeconds) * time.Millisecond):
					c <- fmt.Sprintf("message from %s: %d", name, i)
				case <-done:
					fmt.Println("cleaning up")
					if cleanupConsuming > 0 {
						//清理花费秒数
						time.Sleep(time.Duration(cleanupConsuming) * time.Second)
						fmt.Println("cleaning up finished")
					}
					done <- struct{}{}
					return
				}
			} else {
				time.Sleep(time.Duration(rand.Intn(1000)*tickSeconds) * time.Millisecond)
				c <- fmt.Sprintf("message from %s: %d", name, i)
			}
			i++
		}
	}()
	return c
}

//GetNum get number from str
// message from service1: 12
func GetNum(s string) (num int64) {
	re := regexp.MustCompile(`(?:\w+\s)+\w+:\s(\d+)`)
	numStr := re.FindStringSubmatch(s)
	if len(numStr) > 1 && strings.TrimSpace(numStr[1]) != "" {
		num, _ = strconv.ParseInt(strings.TrimSpace(numStr[1]), 10, 64)
	}
	return

}
