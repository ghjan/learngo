package util

import (
	"errors"
	"strings"
	"time"
)

func ParseTime(timestamp string) (result time.Time, err error) {
	layout := "2006-01-02 15:04:05"
	if strings.Contains(timestamp, ".") &&
		(strings.Contains(timestamp, "-") || strings.Contains(timestamp, "/")) {
		timestamp = timestamp[:strings.LastIndex(timestamp, ".")]
		result, err = time.Parse(layout, timestamp)
	} else {
		err = errors.New("not a valid format time ")
	}
	return
}

//IsNotTimeout:没有超时
func IsNotTimeout(timestamp string, timeout int64) (valid bool) {
	var (
		result time.Time
		err    error
	)
	result, err = ParseTime(timestamp)
	tNow := time.Now()
	if err == nil {
		valid = result.Sub(tNow).Seconds() <= float64(timeout)
	} else {
		valid = ParseFloat64(timestamp)/1000 <= float64(tNow.Unix()+timeout)
	}
	return
}
