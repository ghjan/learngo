package util

import (
	"container/list"
	"github.com/axgle/mahonia"
	mapset "github.com/deckarep/golang-set"
	"github.com/tidwall/gjson"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func PathExists2(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func ConvertToSlice(listStr *list.List) []string {
	sli := []string{}
	for el := listStr.Front(); nil != el; el = el.Next() {
		sli = append(sli, el.Value.(string))
	}

	return sli
}

func IsNum(str string) bool {
	pattern := `^[+-]?\d+.*\d+$`
	//反斜杠要转义
	result, _ := regexp.MatchString(pattern, str)
	return result
}

func IsBool(str string) bool {
	pattern := `^false|true$`
	//反斜杠要转义
	result, _ := regexp.MatchString(pattern, strings.ToLower(strings.TrimSpace(str)))
	return result
}

/*
 * 删除Slice中的元素。
 * params:
 *   s: slice对象，类型为[]interface{}
 *   index: 要删除元素的索引
 * return:
 *   已经删除指定元素的slice，类型为[]interface{}
 * 说明：返回的序列与传入的序列地址不发生变化(但是传入的序列内容已经被修改，不能再使用)
 */
func SliceRemove(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// Float64toStr 64位浮点数转换为字符串
func Float64toStr(inputNum float64, prec int) string {
	// to convert a float64 number to a string
	return strconv.FormatFloat(inputNum, 'f', prec, 64)
}

// Float64toStrWithRound45 64位浮点数四舍五入以后转换为字符串
func Float64toStrWithRound45(inputNum float64, prec int) string {
	// to convert a float64 number to a string
	return strconv.FormatFloat(Round45(inputNum, prec), 'f', prec, 64)
}

// Float32toStr 32位浮点数转换为字符串
func Float32toStr(inputNum float32, prec int) string {
	// to convert a float32 number to a string
	return Float64toStr(float64(inputNum), prec)
}

func ParseFloat64(f string) float64 {
	s, _ := strconv.ParseFloat(f, 64)
	return s
}

func RoundUp(v0 float64, v1 int) float64 {
	v := v0
	if math.Abs(float64(v1)) < 20 {
		c := v0 * math.Pow(10, float64(v1))
		if v1 > 0 && c-float64(int(c)) > 0 && c-float64(int(c)) < 0.0000001 {
			c = float64(int(c))
		}
		v = math.Ceil(c) / math.Pow(10, float64(v1))
	}
	return v

}

func Round45(v0 float64, v1 int) float64 {
	v := v0
	if math.Abs(float64(v1)) < 20 {
		v = math.Trunc(v0*math.Pow(10, float64(v1))+0.5) / math.Pow(10, float64(v1))
	}
	return v
}

func GetNewKeyAndNewParam(param string, old string, new string) (newKey, paramRoom string) {
	//count_room_exist_network_line ->roomExistMap   room_exist_network_line
	newKey = strings.Replace(strings.TrimPrefix(strings.TrimPrefix(param, "count_"), "layout_"),
		"room", "house", 1)
	newKeyParts := strings.Split(newKey, "_")
	newKey = newKeyParts[0] + strings.Title(newKeyParts[1]) + "Map"
	if strings.HasPrefix(param, "layout_") {
		newKey = "layout_" + newKey
	}
	paramRoom = strings.Replace(param, old, new, 1)
	return
}

func AddDigitPart(str string, digitPart string, sep string) string {
	str = strings.TrimSpace(str)
	namePart := GetNamePartFromKey(str, sep)
	if strings.HasSuffix(namePart, digitPart) {
		return namePart
	} else if str == "" {
		return digitPart
	}
	return strings.Join([]string{
		namePart, digitPart,
	}, sep)
}

func GetNamePartFromKey(str string, sep string) string {
	namePart := strings.TrimSpace(str)
	if sep == "" {
		sep = ","
	}
	keys := strings.Split(str, sep)
	if len(keys) > 1 {
		for _, k := range keys {
			if !IsNum(k) {
				namePart = k
				break
			}
		}
	}
	return namePart
}

func GetDigitPartFromKey(key string) string {
	digitPart := key
	keys := strings.Split(key, ",")
	if len(keys) > 1 {
		for _, k := range keys {
			if IsNum(k) {
				digitPart = k
				break
			}
		}
	}
	return digitPart
}

func _getColumnIndex(titles []gjson.Result, search string) (index int) {
	for i, v := range titles {
		if v.String() == search {
			index = i
			break
		}
	}
	return
}

func FixSize(maxSize, minSize string) (needChange bool, maxResult, minResult string) {
	var (
		fmaxSize float64
		fminSize float64
		err      error
	)
	needChange = false
	//处理大小相反的
	if fmaxSize, err = strconv.ParseFloat(maxSize, 64); err == nil {
		fminSize, err = strconv.ParseFloat(minSize, 64)
		if fmaxSize < fminSize {
			needChange = true
			maxResult = minSize
			minResult = maxSize
			return
		}
	}

	return
}

func PercentToFloat(value string) (waste float64, err error) {
	if value != "" && value != "0.00%" {
		hasPercent := strings.Contains(value, "%")
		waste, err = strconv.ParseFloat(strings.TrimSuffix(value, "%"), 64)
		if err == nil && ((hasPercent && waste > 0) || waste >= 1) {
			waste = waste / 100.0
		}
	}
	return waste, err
}

func GetInterfaceSliceFromStringSlice(strSlice []string) (interfaceSlice []interface{}) {
	interfaceSlice = make([]interface{}, len(strSlice))
	for i, d := range strSlice {
		interfaceSlice[i] = d
	}
	return
}

func GetInterfaceSliceFromIntSlice(intSlice []int) (interfaceSlice []interface{}) {
	interfaceSlice = make([]interface{}, len(intSlice))
	for i, d := range intSlice {
		interfaceSlice[i] = d
	}
	return
}

func NewSetFromStringSlice(strSlice []string) (set mapset.Set) {
	wantedInterfaceSlice := GetInterfaceSliceFromStringSlice(strSlice)
	set = mapset.NewSetFromSlice(wantedInterfaceSlice)
	return
}

func NewSetFromIntSlice(intSlice []int) (set mapset.Set) {
	wantedInterfaceSlice := GetInterfaceSliceFromIntSlice(intSlice)
	set = mapset.NewSetFromSlice(wantedInterfaceSlice)
	return
}
