package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetSqlSortString(t *testing.T) {
	var tests = []struct {
		Ordering  string
		WantedSql string
	}{
		{Ordering: "  ", WantedSql: ""},
		{Ordering: " , ", WantedSql: ""},
		{Ordering: "copy_from_id,updated_at DESC", WantedSql: "copy_from_id ASC,updated_at DESC"},
		{Ordering: "-copy_from_id,-updated_at", WantedSql: "copy_from_id DESC,updated_at DESC"},
		{Ordering: "copy_from_id,-updated_at", WantedSql: "copy_from_id ASC,updated_at DESC"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			t.Logf("index:%d\n", index)
			sql := GetSqlSortString(tt.Ordering)

			assert.Equal(t, strings.ToLower(tt.WantedSql), strings.ToLower(sql))
		})
	}

}
