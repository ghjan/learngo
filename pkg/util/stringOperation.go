package util

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// FindNumOrLowerLetter 寻找数字和小写字符
func FindNumOrLowerLetter(str string) []string {
	reg := regexp.MustCompile("[\\d|a-z]+")
	return reg.FindAllString(str, -1)
	//[00az abc09 ab 99]
}

// FindNum 搜索数字
func FindNum(str string) []string {
	reg := regexp.MustCompile("\\d+")
	return reg.FindAllString(str, -1)
	//[00az abc09 ab 99]
}

// RemovePrefixArray 去掉某些前缀
func RemovePrefixArray(str string, prefixArray []string) string {
	for _, prefix := range prefixArray {
		index := strings.Index(str, prefix)
		if index >= 0 {
			return str[index:]
		}
	}
	return str
}

// RemoveNum 去掉包含的数字
func RemoveNum(str string) string {
	reg := regexp.MustCompile("\\d")
	return reg.ReplaceAllString(str, "")
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Camel2Case 驼峰写法转为下划线写法
func Camel2Case(name string) string {
	var isUpperChar, isUpperLastChar bool
	result := ""
	for index, c := range name {
		isUpperChar = unicode.IsUpper(c)
		needInsert := isUpperChar && !isUpperLastChar
		if needInsert {
			if index > 0 {
				result += "_"
			}
			result += string(unicode.ToLower(c))
		} else {
			if isUpperChar {
				result += string(unicode.ToLower(c))
			} else {
				result += string(c)
			}
		}
		isUpperLastChar = isUpperChar
	}
	return result
}

// StrFirstToUpper 首字母大写
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= - 32
	}
	return string(strArry)
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// TrimLeftSpace returns a slice of the string s, with only leading
// white space removed, as defined by Unicode.
func TrimLeftSpace(s string) string {
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			return strings.TrimFunc(s[start:], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// At this point s[start:stop] starts and ends with an ASCII
	// non-space bytes, so we're done. Non-ASCII cases have already
	// been handled above.
	return s[start:]
}

// AppendPrefixIfNotExist 添加前缀字符串（如果前缀不是这个-避免重复添加前缀）
func AppendPrefixIfNotExist(content, prefix string) (result string) {
	result = TrimLeftSpace(content)
	if !strings.HasPrefix(result, prefix) {
		result = prefix + result
	}
	return
}

func AppendNewComerToContent(content string, newComer string) (result string) {
	newComer = strings.TrimSpace(newComer)
	result = strings.TrimSpace(content)
	if newComer == "" {
		return
	}
	if result == "" {
		result = newComer
	} else {
		members := strings.Split(result, ",")
		set := NewSetFromStringSlice(members)
		if set.Cardinality() == 0 || !set.Contains(newComer) {
			result += "," + newComer
		}
	}
	return
}

func ContainsAll(content string, multipleKey string) (result bool) {
	keys := strings.Split(multipleKey, ";")
	for _, k := range keys {
		if !strings.Contains(content, strings.TrimSpace(k)) {
			result = false
			return
		}
	}
	result = true
	return
}

func RemoveSuffixPart(value string, sep string) (result string) {
	if sep == "" {
		sep = "-"
	}
	index := strings.Index(value, "-")
	if index >= 0 {
		result = value[:strings.Index(value, "-")]
	} else {
		result = value
	}
	return
}

func ReplaceModelName(sourceName, sourceUUID, targetUUID string) (targetName string) {
	targetName = strings.Replace(sourceName, sourceUUID, targetUUID, -1)
	if !strings.Contains(targetName, targetUUID) {
		lastIndex := strings.LastIndex(targetName, "_")
		if lastIndex >= 0 {
			targetName = targetName[:lastIndex] + targetUUID
		} else {
			targetName += "_" + targetUUID
		}
	}
	return
}

//ExtractString 获取某个字符串
func ExtractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}

var totalRe = regexp.MustCompile(
	`total\.(.+)`)

func GetPostvarForTotal(contents []byte) string {
	return ExtractString(contents, totalRe)
}

var subitemnoRe = regexp.MustCompile(
	`(\d+)\-\d+`)

func GetItemNoFromSubItemNo(contents []byte) string {
	return ExtractString(contents, subitemnoRe)
}

var typeRe = regexp.MustCompile(
	`型号[:：](.+)\n`)

var specsRe = regexp.MustCompile(
	`规格[:：](.+)\n`)

var descriptionRe = regexp.MustCompile(
	`说明[:：]([\s\S]+)`)

// GetTypeFromSpecification 获取型号
func GetTypeFromSpecification(contents []byte) string {
	return strings.TrimSpace(ExtractString(contents, typeRe))
}

// GetSpecsFromSpecification 获取规格
func GetSpecsFromSpecification(contents []byte) string {
	return strings.TrimSpace(ExtractString(contents, specsRe))
}

// GetDescriptionFromSpecification 获取说明
func GetDescriptionFromSpecification(contents []byte, needFormat bool) (description string) {
	description = strings.TrimSpace(ExtractString(contents, descriptionRe))
	if needFormat {
		description = FormatMultilineString(description, false)
	}
	return
}

func FormatMultilineString(strSource string, removeEmptyNewline bool) (result string) {
	if removeEmptyNewline {
		strSlice := splitByEmptyNewline(strSource)
		for _, str := range strSlice {
			str = strings.TrimSpace(str)
		}
		result = strings.Join(strSlice, "\n")
	} else {
		result = regexp.
			MustCompile("\r\n").
			ReplaceAllString(strSource, "\n")
	}
	return
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

var ossUrlRe = regexp.MustCompile(
	`llib/(.+)/[\s\S]+`)

// GetOssIdFromOssUrl 获取OssId
func GetOssIdFromOssUrl(contents []byte) (ossId string) {
	ossId = strings.TrimSpace(ExtractString(contents, ossUrlRe))
	return
}
