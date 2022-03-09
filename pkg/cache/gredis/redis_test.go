package gredis

import (
	"fmt"
	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLikeDeletes(t *testing.T) {
	var (
		err error
	)
	patternPart := "test_like_deletes"
	var keys []string
	for i := 0; i < 10; i++ {
		key := shortuuid.New() + patternPart + shortuuid.New()
		keys = append(keys, key)
		Set(key, 1, TimeoutNonce)
	}
	for _, key := range keys {
		assert.True(t, Exists(key), fmt.Sprintf("key should be existed:%s", key))
	}
	err = LikeDeletes(patternPart)
	assert.True(t, err == nil, fmt.Sprintf("LikeDeletes expected:err==nil,patternPart:%s", patternPart))
	for _, key := range keys {
		assert.True(t, !Exists(key), fmt.Sprintf("LikeDeletes expected:!Exists(key:%s)", key))
	}
	err = LikeDeletes(patternPart)
	assert.True(t, err == nil, fmt.Sprintf("LikeDeletes expected:err==nil,patternPart:%s", patternPart))
}

//
//func TestMain(m *testing.M) {
//	fmt.Println("begin")
//	//db.InitDB()
//
//	m.Run()
//	fmt.Println("end")
//}
