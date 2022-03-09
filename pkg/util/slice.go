package util

import (
	"errors"
	"github.com/tidwall/gjson"
	"reflect"
)

// StructSliceToMapSlice : struct切片转为map切片
func StructSliceToMapSlice(source interface{}) (res []map[string]interface{}, err error) {
	// 判断，interface转为[]interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		err = errors.New("ERROR: Unknown type, slice expected.")
		return
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	// 通过遍历，每次迭代将struct转为map
	for _, elem := range ret {
		data := make(map[string]interface{})
		objT := reflect.TypeOf(elem)
		objV := reflect.ValueOf(elem)
		for i := 0; i < objT.NumField(); i++ {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
		res = append(res, data)
	}
	return
}

// StructSliceToMap : struct切片转为map
func StructSliceToMap(source interface{}, keyFieldName string) (res map[string]interface{}, err error) {
	// 判断，interface转为[]interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		err = errors.New("ERROR: Unknown type, slice expected.")
		return
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	// 通过遍历，每次迭代将struct转为map
	res = make(map[string]interface{})
	for index, elem := range ret {
		data := make(map[string]interface{})
		objT := reflect.TypeOf(elem)
		objV := reflect.ValueOf(elem)
		for i := 0; i < objT.NumField(); i++ {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
		var key string
		if keyFieldName != "" {
			key = data[keyFieldName].(string)
			res[key] = data
		} else { //keyFieldName为空格的情况下面,返回类型是 map[string]int
			if objT.String() == "gjson.Result" {
				key = elem.(gjson.Result).String()
			} else {
				key = elem.(string)
			}
			res[key] = index
		}
	}
	return
}
