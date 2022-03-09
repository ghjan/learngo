package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ghjan/learngo/pkg/utils"
	"hash"
	"hash/crc32"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Update(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func Sha256(data []byte) string {
	_sha2 := sha256.New()
	_sha2.Write(data)
	return hex.EncodeToString(_sha2.Sum(nil))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func FileSha1Withpath(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("FileSha1_withpath err:%s", err.Error())
		return ""
	}
	return FileSha1(file)
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func StringMD5(source string) string {
	_md5 := md5.New()
	io.WriteString(_md5, source)
	return hex.EncodeToString(_md5.Sum(nil))
}

func String2Map(info string) (result map[string]interface{}, err error) {
	if info == "" {
		return
	}
	//1.创建json数据
	b := []byte(info)
	//2.声明空interface
	var i interface{}
	//3.解析
	if err = json.Unmarshal(b, &i); err != nil || i == nil {
		utils.Checkerr(err, fmt.Sprintf("json.Unmarshal error, i:%#v", i))
	} else {
		//默认转成了map
		//4.解析到interface的json可以判断类型
		result = i.(map[string]interface{})
	}
	return
}

func String2StringMap(info string) (result map[string]string, err error) {
	//1.创建json数据
	b := []byte(info)
	//2.声明空interface
	var i interface{}
	//3.解析
	err = json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	//默认转成了map
	//4.解析到interface的json可以判断类型
	r := i.(map[string]interface{})
	result = make(map[string]string, len(r))
	for k, v := range r {
		switch vv := v.(type) {
		case string:
			result[k] = vv
		case float64:
			result[k] = Float64toStr(vv, 3)
		case float32:
			result[k] = Float32toStr(vv, 3)
		case int:
			result[k] = strconv.Itoa(vv)
		default:
			result[k] = ""
			if err == nil {
				err = errors.New(fmt.Sprintf("unknown type of v:%v", v))
			}
		}
	}
	return
}

func GetMapSortedKeys(m map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

//获取map的排序过的key切片
func GetStringMapSortedKeys(m map[string]string) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// HashCode hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func HashCodeStr(s string) string {
	return strconv.Itoa(HashCode(s))
}

func GetAllVar(str string) (matched []string) {
	pattern := `(?U)\[([A-Za-z0-9_.]+)\]`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(str, -1)
	for _, m := range matches {
		matched = append(matched, strings.TrimSuffix(strings.TrimPrefix(m, "["), "]"))
	}
	return
}
