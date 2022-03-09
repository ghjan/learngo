package util

import (
	"bufio"
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"math"
	"os"
)

//	读取文件中的数据
func ReadJsonFile(filePath string) (result string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	defer file.Close()
	if err != nil {
		return
	}
	buf := bufio.NewReader(file)
	for {
		var s string
		s, err = buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return
			}
		}
	}
	return
}

//ReadJsonFileWithSearch	读取文件中的数据
func ReadJsonFileWithSearch(filePath string, searchDepth int) (result string, err error) {
	filePath1 := filePath
	exists := false
	for i := 0; i < searchDepth; i++ {
		if exists = PathExists(filePath1); !exists {
			filePath1 = "../" + filePath1
			continue
		}
	}
	if exists {
		return ReadJsonFile(filePath1)
	} else {
		return "", errors.New("not existed")
	}
}

func GetStamp(jsonString string) string {
	designStamp := ParseFloat64(gjson.Get(jsonString, "data.designStamp").String())
	gpStamp := ParseFloat64(gjson.Get(jsonString, "data.manual.global_property.timestamp").String())
	profileTimestamp := ParseFloat64(gjson.Get(jsonString, "data.profile.timestamp").String())
	stamp := Float64toStr(math.Max(profileTimestamp, math.Max(designStamp, gpStamp)), 0)
	return stamp
}
