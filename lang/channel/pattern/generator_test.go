package pattern

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsgGen(t *testing.T) {
	m1 := MsgGen("service1", 2, nil, 2)
	m2 := MsgGen("service2", 2, nil, 2)
	for {
		mes1 := <-m1
		fmt.Println(mes1)
		n := GetNum(mes1)
		mes2 := <-m2
		fmt.Println(mes2)
		m := GetNum(mes2)
		if n > 10 || m > 10 {
			break
		}
	}
}

func TestGetNum(t *testing.T) {
	num := GetNum("message from service1: 12")
	assert.Equal(t, int64(12), num)
	num = GetNum("message from micro service1: 12")
	assert.Equal(t, int64(12), num)
}
