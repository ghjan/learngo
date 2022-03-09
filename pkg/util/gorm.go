package util

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"strconv"
)

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

func NewJSON(info string) JSON {
	result := []byte(info)
	return result
}

func JsonMapEqual(quantValues1 JSON, quantValues2 string) bool {
	qv, _ := quantValues1.MarshalJSON()
	quantValuesMap, _ := String2Map(string(qv))
	quantValuesMapSearch, _ := String2Map(quantValues2)
	result := false
	for k, v := range quantValuesMapSearch {
		result = false
		if quantValuesMap[k] != v {
			switch vv := quantValuesMap[k].(type) {
			case float64:
				switch vv2 := v.(type) {
				case string:
					vv2F64, err := strconv.ParseFloat(vv2, 64)
					if err == nil {
						result = vv == vv2F64
					}
				}
			}
			if !result {
				return result
			}
		}
	}
	return true
}

func FindRightPath(filePath string, tryTimes int) (realPath string, err error) {
	realPath = filePath
	for i := 0; i < tryTimes; i++ {
		existed := false
		if existed = PathExists(realPath); existed {
			return
		} else {
			realPath = "../" + realPath
		}
	}
	return "", errors.New("not found")
}
