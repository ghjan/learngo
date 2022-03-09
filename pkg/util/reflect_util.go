package util

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func ReflectSetFieldValue(o interface{}, fieldName string, value interface{}) {
	pp := reflect.ValueOf(o) // 取得struct变量的指针
	fieldName = Case2Camel(fieldName)
	field := pp.Elem().FieldByName(fieldName) //获取指定Field  pp.Elem().FieldByName(fieldName)

	// 设置值
	k := field.Kind()
	switch k {
	case reflect.String:
		field.SetString(value.(string))
	case reflect.Int:
		field.SetInt(int64(value.(int)))
	case reflect.Bool:
		field.SetBool(value.(bool))
	case reflect.Float64:
		field.SetFloat(value.(float64))
	case reflect.Float32:
		field.SetFloat(float64(value.(float32)))
	}
}

func ReflectGetFieldValue(o interface{}, fieldName string) (value interface{}) {
	v := reflect.ValueOf(o)
	fieldName = Case2Camel(fieldName)
	value = v.Elem().FieldByName(fieldName).Interface()
	return
}

func ReflectGetMapFromObject(data interface{}, rf url.Values,
	excludeFieldNames string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	fieldNames, fieldValues := GetAllFieldsOfStruct(data)
	excludeFieldNameSlice := strings.Split(excludeFieldNames, ",")
	excludeFieldNameSet := NewSetFromStringSlice(excludeFieldNameSlice)
	for i, fieldName := range fieldNames {
		fieldNameCase := Camel2Case(fieldName)
		if excludeFieldNameSet != nil && excludeFieldNameSet.Cardinality() > 0 &&
			(excludeFieldNameSet.Contains(fieldName) || excludeFieldNameSet.Contains(fieldNameCase)) {
			continue
		}
		fieldValue := fieldValues[i]
		isBlank := InterfaceIsNil(fieldValue)
		if isBlank && rf != nil && !rf.Has(fieldNameCase) {
			//只处理特定的那些字段的空白，这样就不会错误的把没有提交的那些字段设置为空白
			continue
		}
		result[fieldName] = fieldValue
	}
	return
}

func GetAllFieldsOfStruct(q interface{}) (fieldNames []string, fieldValues []interface{}) {
	value := reflect.ValueOf(q)
	if reflect.Struct != value.Kind() {
		fmt.Printf("unsupported type:%s\n", value.Kind())
		return
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldName := field.Name
		fieldNames = append(fieldNames, fieldName)
		fieldValue := value.FieldByName(fieldName).Interface()
		fieldValues = append(fieldValues, fieldValue)
	}
	return
}
