package remote_sign

import (
	"github.com/ghjan/learngo/pkg/util"
	"sort"
)

//获取map类型排序的keys
func GetSortedKeysOfMap(dictSource map[string]string) (keys []string) {
	keys = make([]string, 0)
	for k, _ := range dictSource {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	return
}

//计算签名
func CalcSignMd5(paramsDict map[string]string, secretKey string) (sign string) {
	keys := GetSortedKeysOfMap(paramsDict)
	string := ""
	for _, key := range keys {
		if key != "sign" && paramsDict[key] != "" {
			string += key + "=" + paramsDict[key] + "&"
		}
	}
	if secretKey != "" {
		string += secretKey
	}

	sign = util.StringMD5(string)

	return
}
