package gredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHMOperations(t *testing.T) {
	var (
		err error
	)
	c := GetRedisConn()
	defer c.Close()
	kvs := map[string]string{
		"bike2": "bluegogo",
		"bike3": "xiaoming",
		"bike4": "xiaolan",
	}
	args := []interface{}{
		"myhash",
	}
	args2 := []interface{}{
		"myhash",
	}
	values := []string{
	}
	for k, v := range kvs {
		values = append(values, k)
		args2 = append(args2, k)
		args = append(args, k)
		args = append(args, v)
	}
	_, err = c.Do("hmset", args...)
	if err != nil {
		t.Logf("hmset error:%s", err.Error())
		t.FailNow()
	} else {
		value, err := redis.Values(c.Do("hmget", args2...))
		if err != nil {
			t.Errorf("hmget failed:%s", err.Error())
			t.FailNow()
		} else {
			for index, v := range value {
				key := values[index]
				if expectedV, existed := kvs[key]; existed {
					assert.Equal(t, expectedV, string(v.([]byte)))
				} else {
					assert.Nil(t, v)
				}
			}
		}
	}
}

func TestHMSet(t *testing.T) {
	var (
		err error
	)
	name := "myhash"
	kvs := map[string]string{
		"bike2": "bluegogo",
		"bike3": "xiaoming",
		"bike4": "xiaolan",
	}
	if err = HMSet(name, kvs); err != nil {
		t.Logf("TestHMSet, HMSet error:%s", err.Error())
		t.FailNow()
	}
}
func TestHMGet(t *testing.T) {
	var (
		value []interface{}
		err   error
	)
	values := []string{
		"1", "2",
	}

	if value, err = HMGet("js_d16a2d82ac903e1d03aa46abb4787a21", values); err != nil {
		// handle error
		t.Errorf("TestMGet, HMGet error:%s\n", err.Error())
		t.FailNow()
	}
	for index, v := range value {
		key := values[index]
		t.Logf("%s:%s", key, string(v.([]byte)))
	}
}

func TestHGet(t *testing.T) {
	var (
		v   interface{}
		err error
	)
	key := "1"
	if v, err = HGet("js_d16a2d82ac903e1d03aa46abb4787a21", key); err != nil {
		// handle error
		t.Errorf("TestMGet, HMGet error:%s\n", err.Error())
		t.FailNow()
	}
	value := string(v.([]byte))
	t.Logf("%s:%s", key, value)
	assert.Contains(t, value, ".json")
}

func TestHExists(t *testing.T) {
	tests := []struct {
		key             string
		expectedExisted bool
	}{
		{
			key: "1", expectedExisted: true,
		},
		{
			key: "501", expectedExisted: false,
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprint(test), func(t *testing.T) {
			var (
				isExist bool
				err     error
			)

			t.Logf("index:%d\n", index)
			if isExist, err = HExists("js_d16a2d82ac903e1d03aa46abb4787a21", test.key); err != nil {
				assert.Errorf(t, err, "TestHexists, HExists error")
				t.FailNow()
			} else {
				assert.Equal(t, test.expectedExisted, isExist)
			}

		})
	}
}

func TestHLen(t *testing.T) {
	var (
		length int64
		err    error
	)
	if length, err = HLen("js_d16a2d82ac903e1d03aa46abb4787a21"); err != nil {
		assert.Errorf(t, err, "TestHLen, HLen error")
		t.FailNow()
	} else {
		t.Logf("length:%d", length)
		assert.True(t, length > 100)
	}
}

func TestHKeys(t *testing.T) {
	var (
		keys []string
		err  error
	)
	if keys, err = HKeys("js_d16a2d82ac903e1d03aa46abb4787a21"); err != nil {
		assert.Errorf(t, err, "TestHKeys, HKeys error")
		t.FailNow()
	} else {
		t.Logf("len(keys):%d,keys[0]:%s", len(keys), keys[0])
		assert.True(t, keys != nil && len(keys) > 100)
	}

}

func TestHVals(t *testing.T) {
	var (
		values []string
		err    error
	)
	if values, err = HVals("js_d16a2d82ac903e1d03aa46abb4787a21"); err != nil {
		assert.Errorf(t, err, "TestHVals, HVals error")
		t.FailNow()
	} else {
		t.Logf("len(values):%d,values[0]:%s", len(values), values[0])
		assert.True(t, values != nil && len(values) > 100)
	}

}

func TestHDel(t *testing.T) {
	var (
		err error
	)
	if err = HDel("myhash", "bike4"); err != nil {
		assert.Errorf(t, err, "TestHDel, HDel error")
		t.FailNow()
	}
}

func TestHGetAll(t *testing.T) {
	var (
		kvs map[string]string
		err error
	)
	if kvs, err = HGetAll("myhash"); err != nil {
		assert.Errorf(t, err, "TestHGetAll, HGetAll error")
		t.FailNow()
	} else {
		t.Logf("kvs:%#v", kvs)
	}
}
