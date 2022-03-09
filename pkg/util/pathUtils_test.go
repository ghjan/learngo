package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsModelDir(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{"1-吉舍特价砖300300砖3110A地砖68元平米", true},
		{"15-吉舍特价砖300600砖6508B墙砖68元平米", true},
		{"瓷砖", false},
		{"1-卡诺亚衣柜Y101实木颗粒板自购", true},
		{"6-史密斯热水器CEWH-60BP电自购", true},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.arg), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := IsModelDir(tt.arg)
			assert.Equal(t, tt.want, result, "TestIsModelDir error")
		})
	}

}
