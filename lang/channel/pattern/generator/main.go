package main

import (
	"fmt"
	"github.com/ghjan/learngo/lang/channel/pattern"
)

func main() {
	m := pattern.MsgGen("", 2, nil, 2)
	for {
		fmt.Println(<-m)
	}
}
