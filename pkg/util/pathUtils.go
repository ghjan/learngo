package util

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func RootPath() (dir string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	return
}

func IsModelDir(str string) bool {
	pattern := "^\\d+\\-.+$"
	//反斜杠要转义
	result, _ := regexp.MatchString(pattern, str)
	return result
}
