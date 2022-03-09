package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFloat64toStr(t *testing.T) {
	a := assert.New(t)
	a.Equal("21312421.21", Float64toStr(21312421.213123, 2),
		" what I expect is 21312421.21 in string format")
	a.Equal("21312421.22", Float64toStr(21312421.215123, 2),
		" what I expect is 21312421.22 in string format")
	a.Equal("8.22", Float64toStr(8.225, 2),
		" what I expect is 8.22 in string format")
}

type TestRoundUpArgs struct {
	Args []string
	Want float64
}

var testRoundUpArgs = []TestRoundUpArgs{
	{Args: []string{"34.77", "-1"}, Want: 40},
	{Args: []string{"34.77", "0"}, Want: 35},
	{Args: []string{"34.76", "2"}, Want: 34.76},
	{Args: []string{"34.77", "2"}, Want: 34.77},
	{Args: []string{"34.768", "2"}, Want: 34.77},
	{Args: []string{"34.763", "2"}, Want: 34.77},
	{Args: []string{"34.760", "2"}, Want: 34.76},
	{Args: []string{"34.770", "2"}, Want: 34.77},
	{Args: []string{"21312421.213123", "2"}, Want: 21312421.22},
	{Args: []string{"21312421.215123", "2"}, Want: 21312421.22},
	{Args: []string{"8.225", "2"}, Want: 8.23},
}

func TestRoundUp(t *testing.T) {
	for index, tt := range testRoundUpArgs {
		//a := assert.New(t)
		t.Run(fmt.Sprint(tt.Args[0]+"-"+tt.Args[1]), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			//msg := fmt.Sprintf(" what I expect is %f", tt.Want)
			//a.Equal(tt.Want, RoundUp(ParseFloat64(tt.Args[0]), int(ParseFloat64(tt.Args[1]))), msg)
			v0 := ParseFloat64(tt.Args[0])
			v1 := int(ParseFloat64(tt.Args[1]))
			result := RoundUp(v0, v1)
			assert.Equal(t, tt.Want, result)
			if tt.Want != result {
				tmp := float64(int(v0 * math.Pow(10, float64(v1))))
				t.Log(tmp)
				t.Log(math.Ceil(tmp))
				t.Fail()
			}
		})
	}
}

func TestRound45(t *testing.T) {
	a := assert.New(t)
	a.Equal(21312421.21, Round45(21312421.213123, 2),
		" what I expect is 21312421.21")
	a.Equal(21312421.22, Round45(21312421.215123, 2),
		" what I expect is 21312421.22")
	a.Equal(8.22, Round45(8.224, 2),
		" what I expect is 8.22")
	a.Equal(8.23, Round45(8.225, 2),
		" what I expect is 8.23")
	a.Equal(4550.00, Round45(4545, -1),
		" what I expect is 4550")
}

func TestFloat64toStrWithRound45(t *testing.T) {
	a := assert.New(t)
	a.Equal("21312421.21", Float64toStrWithRound45(21312421.213123, 2),
		" what I expect is 21312421.21 in string format")
	a.Equal("21312421.22", Float64toStrWithRound45(21312421.215123, 2),
		" what I expect is 21312421.22 in string format")
	a.Equal("8.23", Float64toStrWithRound45(8.225, 2),
		" what I expect is 8.23 in string format")
	a.Equal("8.22", Float64toStrWithRound45(8.224, 2),
		" what I expect is 8.22 in string format")
}

func TestGetNamePartFromKey(t *testing.T) {
	tests := []struct {
		arg  string
		sep  string
		want string
	}{
		//{"105,卧室1", ",", "卧室1"},
		//{"105,卫生间", ",", "卫生间"},
		//{"brand_201", "_", "brand"},
		//{"brand", "_", "brand"},
		//{"zigou_806", "_", "zigou"},
		{"806", ",", "806"},
		//{"806", "_", ""},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.arg), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := GetNamePartFromKey(tt.arg, tt.sep)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestAddDigitPart(t *testing.T) {
	tests := []struct {
		arg       string
		digitPart string
		sep       string
		want      string
	}{
		{"brand_201", "201", "_", "brand_201"},
		{"zigou", "806", "_", "zigou_806"},
		{"zigou_806", "806", "_", "zigou_806"},
		{"806", "806", "_", "806"},
		{"", "806", "_", "_806"},
		//{"806", "_", ""},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.arg), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := AddDigitPart(tt.arg, tt.digitPart, tt.sep)
			assert.Equal(t, tt.want, result)
		})
	}

}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{"105", true},
		{"6.25", true},
		{"+6.25", true},
		{"-6.25", true},
		{"卧室1", false},
		{"卫生间", false},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.arg), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := IsNum(tt.arg)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIsBool(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{"true", true},
		{"false", true},
		{"TRUE", true},
		{"FALSE", true},
		{" TRUE ", true},
		{" FALSE ", true},
		{"105", false},
		{"6.25", false},
		{"+6.25", false},
		{"-6.25", false},
		{"卧室1", false},
		{"卫生间", false},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt.arg), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			result := IsBool(tt.arg)
			assert.Equal(t, tt.want, result)
		})
	}
}
