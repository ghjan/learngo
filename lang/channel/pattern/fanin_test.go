package pattern

import (
	"fmt"
	"testing"
)

func TestFanInByGoroutine(t *testing.T) {
	m1 := MsgGen("service1", 2, nil, 2)
	m2 := MsgGen("service2", 2, nil, 2)
	m3 := MsgGen("service3", 2, nil, 2)
	m := FanInByGoroutine(m1, m2, m3)
	for {
		mes := <-m
		fmt.Println(mes)
		num := GetNum(mes)

		if num > 10 {
			break
		}
	}
}

func TestFanInBySelect(t *testing.T) {
	m1 := MsgGen("service1", 2, nil, 2)
	m2 := MsgGen("service2", 2, nil, 2)
	m := FanInBySelect(m1, m2)
	for {
		mes := <-m
		fmt.Println(mes)
		num := GetNum(mes)

		if num > 10 {
			break
		}
	}
}
