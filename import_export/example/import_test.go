package example

import (
	"encoding/json"
	"os"
	"testing"
)

func TestBulkIndex(t *testing.T) {
	var (
		file         *os.File
		data         []Course
		successItems []Course
		err          error
	)

	if file, err = os.Open("data.json"); err != nil {
		t.Errorf("TestBulkIndex, os.Open error:%s\n", err.Error())
		t.FailNow()
	} else {
		defer file.Close()
	}

	jsonDecoder := json.NewDecoder(file)
	if err = jsonDecoder.Decode(&data); err != nil {
		t.Errorf("TestBulkIndex, Decode error:%s\n", err.Error())
		t.FailNow()
	}

	if successItems, err = BulkIndex("imooc_test", "course", data); err != nil {
		t.Errorf("TestBulkIndex, BulkIndex error:%s\n", err.Error())
		t.FailNow()
	} else {
		t.Logf("successItems:%#v\n", successItems)
	}

}
