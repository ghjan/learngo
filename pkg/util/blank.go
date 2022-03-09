package util

import "reflect"

func InterfaceIsNil(i interface{}) (ret bool) {
	ret = i == nil

	if !ret { //需要进一步做判断
		vi := reflect.ValueOf(i)
		kind := vi.Kind()
		switch kind {
		case reflect.Slice, reflect.Map, reflect.Chan, reflect.Interface, reflect.Func, reflect.Ptr:
			ret = vi.IsNil()
		default:
			ret = vi.IsZero()
		}
	}

	return
}
