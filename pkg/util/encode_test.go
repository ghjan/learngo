package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString2Map(t *testing.T) {
	//a := assert.New(t)
	result, err := String2Map(`{"Name":"zhangsan","Hobby":"女","Age":33}`)
	assert.Equal(t,nil, err)
	for k, v := range result {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "是string类型", vv)
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case float32:
			fmt.Println(k, "是float32类型", vv)
		case int:
			fmt.Println(k, "是int类型", vv)
		default:
			fmt.Println("xxx")
		}
	}
	assert.Equal(t,"zhangsan", result["Name"].(string))
	assert.Equal(t,"女", result["Hobby"].(string))
	assert.Equal(t,33.0, result["Age"].(float64))
}

func TestHashCode(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"201-brand-西门子", "3404487580"},
		{"202-brand-施耐德-leakage_protector-总开带漏报", "645722910"},
		{"202-brand-施耐德-leakage_protector-分路单片带漏保", "3378720709"},
	}
	//a := assert.New(t)
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			code := HashCodeStr(tt.args)
			assert.Equal(t,tt.want, code)
		})
	}
}
