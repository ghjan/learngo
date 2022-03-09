package util

import (
	"net/url"
	"strings"
)

func GetQueryParameterValue(queryString string, key string) (value string, queryKv map[string]string) {
	pares := strings.Split(queryString, "&")
	key = strings.TrimSpace(key)
	queryKv = make(map[string]string, 0)
	for _, pare := range pares {
		kv := strings.Split(pare, "=")
		kv[0] = strings.TrimSpace(kv[0])
		if len(kv) == 2 && kv[0] != "" {
			kv[1], _ = url.QueryUnescape(kv[1])
			queryKv[kv[0]] = kv[1]
			if key != "" && kv[0] == key {
				value = kv[1]
			}
		}
	}
	return
}
