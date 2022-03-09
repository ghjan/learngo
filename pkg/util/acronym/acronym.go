package acronym

import "github.com/axgle/pinyin"

type AcronumMethod func(string) string

var AcronumMethods = []AcronumMethod{
	GetAcronymNeeddigit,
	GetAcronymFullPingying,
	GetAcronymFullPingyingNeeddigit,
}

func GetAcronym(chineseWord string) (acronym string) {
	pyWord := pinyin.Convert(chineseWord)
	for _, a := range pyWord {
		if a >= 'A' && a <= 'Z' {
			acronym += string(a)
		}
	}
	return
}

func GetAcronymFullPingying(chineseWord string) (acronym string) {
	pyWord := pinyin.Convert(chineseWord)
	for _, a := range pyWord {
		if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') {
			acronym += string(a)
		}
	}
	return
}

func GetAcronymNeeddigit(chineseWord string) (acronym string) {
	pyWord := pinyin.Convert(chineseWord)
	for _, a := range pyWord {
		if (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
			acronym += string(a)
		}
	}
	return
}

func GetAcronymFullPingyingNeeddigit(chineseWord string) (acronym string) {
	pyWord := pinyin.Convert(chineseWord)
	for _, a := range pyWord {
		if (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
			acronym += string(a)
		}
	}
	return
}
