package util

import "strings"

func GetSqlSortString(ordering string) (sortString string) {
	orderings := strings.Split(ordering, ",")
	sorts := make([]string, 0)
	for _, order1 := range orderings {
		var (
			fieldName string
		)
		order1 = strings.TrimSpace(strings.ToLower(order1))
		fieldNames := strings.Split(strings.TrimPrefix(strings.TrimPrefix(order1, "-"), "+"), " ")
		if len(fieldNames) > 0 {
			fieldName = strings.TrimSpace(fieldNames[0])
		}
		if fieldName == "" {
			continue
		}
		isDesc := strings.HasPrefix(order1, "-") || strings.HasSuffix(order1, "desc")
		if isDesc {
			sorts = append(sorts, fieldName+" DESC")
		} else {
			sorts = append(sorts, fieldName+" ASC")
		}
	}
	if sorts != nil && len(sorts) > 0 {
		sortString = strings.Join(sorts, ",")
	}
	return
}
