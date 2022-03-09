package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceIsNil(t *testing.T) {
	var tests = []struct {
		Comment        string
		ValueBeChecked interface{}
		WantedBeBlank  bool
	}{
		{Comment: "nil", ValueBeChecked: nil, WantedBeBlank: true},
		{Comment: "字符串空值", ValueBeChecked: "", WantedBeBlank: true},
		{Comment: "bool false", ValueBeChecked: false, WantedBeBlank: true},
		{Comment: "int 0 ", ValueBeChecked: 0, WantedBeBlank: true},
		{Comment: "float64 0 ", ValueBeChecked: 0.0, WantedBeBlank: true},
		{Comment: "切片[]int(nil) ", ValueBeChecked: []int(nil), WantedBeBlank: true},
		{Comment: "映射map[int]int(nil) ", ValueBeChecked: map[int]int(nil), WantedBeBlank: true},
		{Comment: "通道chan int(nil)", ValueBeChecked: chan int(nil), WantedBeBlank: true},
		{Comment: "函数(func())(nil) ", ValueBeChecked: (func())(nil), WantedBeBlank: true},
		{Comment: "指针(*int)(nil)", ValueBeChecked: (*int)(nil), WantedBeBlank: true},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := InterfaceIsNil(tt.ValueBeChecked)
			assert.Equal(t, tt.WantedBeBlank, result, "WantedBeBlank==result")
		})
	}
}
