package utils

import "regexp"

func HasChineseChar(str string) bool {
	var a = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	//接受正则表达式的范围
	for _, v := range str {
		//golang中string的底层是byte类型，所以单纯的for输出中文会出现乱码，这里选择for-range来输出
		if a.MatchString(string(v)) {
			//判断是否为中文，如果是返回一个true，不是返回false。这俩面MatchString的参数要求是string
			//但是 for-range 返回的 value 是 rune 类型，所以需要做一个 string() 转换
			//fmt.Printf("str 字符串第 %v 个字符是中文。是“%v”字\n", i+1, string(v))
			return true
		}
	}
	return false
}
