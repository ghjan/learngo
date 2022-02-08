package main

import (
	"fmt"
	"github.com/ghjan/learngo/lang/channel/pattern"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
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
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m := pattern.FanInBySelect(m1, m2)
	for {
		fmt.Println(<-m)
	}
}
