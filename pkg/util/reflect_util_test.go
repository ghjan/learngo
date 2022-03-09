package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllFieldsOfStruct(t *testing.T) {
	var tests = []struct {
		Name   string
		Age    int
		Salary float64
	}{
		{Name: "person1", Age: 30, Salary: 10000.50},
		{Name: "", Age: 0, Salary: 0},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			fieldNames, fieldValues := GetAllFieldsOfStruct(tt)
			for i, fieldName := range fieldNames {
				t.Logf("fieldName:%s", fieldName)
				fieldValue := fieldValues[i]
				value := ReflectGetFieldValue(&tt, fieldName)
				t.Logf("value:%v,fieldValue:%v", value, fieldValue)
				assert.Equal(t, value, fieldValue)
			}
		})
	}
}
