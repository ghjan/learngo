package utils

import "strings"

var escapeChar = make(map[string]string, 2)
var escapeValuesChar = make(map[string]string, 2)

func init() {
	escapeChar["mysql"] = "`"
	escapeChar["postgres"] = "\""
	escapeValuesChar["mysql"] = "`"
	escapeValuesChar["postgres"] = "'"
}

func EscapeString(driverName, str string) string {
	return escapeChar[driverName] + strings.TrimSpace(str) + escapeChar[driverName]
}

func EscapeValuesString(driverName, str string) string {
	return escapeValuesChar[driverName] + strings.TrimSpace(str) + escapeValuesChar[driverName]
}
func EscapeSpecificChar(str string) string {
	result := ""
	str2 := []rune(strings.TrimSpace(str))
	for i := 0; i < len(str2); i++ {
		if string(str2[i]) == "\\" {
			result += "\\\\"
		} else {
			result += string(str2[i])
		}
	}
	return result
}
