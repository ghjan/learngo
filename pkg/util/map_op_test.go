package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeMap(t *testing.T) {
	var tests = []struct {
		map1           map[string]string
		map2           map[string]string
		expectedResult map[string]string
	}{
		{map1: map[string]string{"1": "110", "2": "120", "3": "119"},
			map2:           map[string]string{"1": "111", "2": "122", "4": "129"},
			expectedResult: map[string]string{"1": "111", "2": "122", "3": "119", "4": "129"},
		},
		{map1: nil,
			map2:           map[string]string{"1": "111", "2": "122", "4": "129"},
			expectedResult: map[string]string{"1": "111", "2": "122", "4": "129"},
		},
		{map1: map[string]string{"1": "110", "2": "120", "3": "119"},
			map2:           nil,
			expectedResult: map[string]string{"1": "110", "2": "120", "3": "119"},
		},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			n := MergeMap(tt.map1, tt.map2)
			assert.Equal(t, tt.expectedResult, n)
			t.Logf("n:%#v", n)
		})
	}

}

func TestGetValueFromMap(t *testing.T) {
	const DEFAULT_MODEL_COMMON = `飘窗:NojScRSE6kzKj9Zyb7rNoa,移门:ENqNLZ97XYkM9epALRomsT,入户门:mE2XPc2Vpu3poMQfwoapXm,
室内门:VSUHWF5hS8LesBSTHCjann,开窗:B2kuvX8afczAvAeqkFpcfF,移窗:NojScRSE6kzKj9Zyb7rNoa,门洞:fRyV5TiYkVMUUBaKJo9bwm,窗洞:7Tbj2MBKovCHf4jJ7F948W`
	tests := []struct {
		k             string
		expectedValue string
	}{
		{k: "移门", expectedValue: "ENqNLZ97XYkM9epALRomsT"},
		{k: "飘窗", expectedValue: "NojScRSE6kzKj9Zyb7rNoa"},
		{k: "入户门", expectedValue: "mE2XPc2Vpu3poMQfwoapXm"},
		{k: "室内门", expectedValue: "VSUHWF5hS8LesBSTHCjann"},
		{k: "开窗", expectedValue: "B2kuvX8afczAvAeqkFpcfF"},
		{k: "移窗", expectedValue: "NojScRSE6kzKj9Zyb7rNoa"},
		{k: "门洞", expectedValue: "fRyV5TiYkVMUUBaKJo9bwm"},
		{k: "窗洞", expectedValue: "7Tbj2MBKovCHf4jJ7F948W"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			var (
				value string
				err   error
			)
			t.Logf("index:%d\n", index)
			value, err = GetValueFromMap(DEFAULT_MODEL_COMMON, tt.k)
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedValue, value)
		})
	}

}
