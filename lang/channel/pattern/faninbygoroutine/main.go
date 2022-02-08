package main

import (
	"fmt"
	"github.com/ghjan/learngo/lang/channel/pattern"
	"math/rand"
	"time"
)

func MsgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("message from %s: %d", name, i)
			i++
		}
	}()
	return c
}

func main() {
	m1 := MsgGen("service1")
	m2 := MsgGen("service2")
	m3 := MsgGen("service3")
	m := pattern.FanInByGoroutine(m1, m2, m3)
	for {
		fmt.Println(<-m)
	}
}
