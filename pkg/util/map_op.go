package util

import "strings"

/*
MergeMap 合并两个词典,并且后面一个词典的覆盖前面一个
*/
func MergeMap(x, y map[string]string) (n map[string]string) {

	if x == nil || len(x) == 0 {
		n = y
		return
	} else if y == nil || len(y) == 0 {
		n = x
		return
	}
	n = make(map[string]string)
	for i, v := range x {
		for j, w := range y {
			if i == j {
				n[i] = w
			} else {
				if _, ok := n[i]; !ok {
					n[i] = v
				}
				if _, ok := n[j]; !ok {
					n[j] = w
				}
			}
		}
	}
	return
}

func GetValueFromMap(kvStr string, key string) (value string, err error) {
	dms := strings.Split(kvStr, ",")
	for _, dm := range dms {
		dm = strings.TrimSpace(dm)
		ds := strings.Split(dm, ":")
		if len(ds) == 2 {
			if ds[0] == key {
				value = ds[1]
				return
			}
		}
	}
	return
}
func GetMapFromListStr(kvStr string) (resultMap map[string]string, err error) {
	resultMap = make(map[string]string, 0)
	dms := strings.Split(kvStr, ",")
	for _, dm := range dms {
		dm = strings.TrimSpace(dm)
		ds := strings.Split(dm, ":")
		if len(ds) == 2 {
			resultMap[ds[0]] = ds[1]
		}
	}
	return
}
