package gredis

import (
	"github.com/gomodule/redigo/redis"
)

func HExists(name string, key string) (isExist bool, err error) {
	conn := GetRedisConn()
	defer conn.Close()
	value, err := conn.Do("hexists", name, key)
	if err == nil {
		isExist = value.(int64) > 0
	}
	return
}

func HLen(name string) (length int64, err error) {
	conn := GetRedisConn()
	defer conn.Close()
	value, err := conn.Do("hlen", name)
	if err == nil {
		length = value.(int64)
	}
	return
}

func HGet(name string, key string) (reply interface{}, err error) {
	conn := GetRedisConn()
	defer conn.Close()
	args := []interface{}{
		name,
		key,
	}
	reply, err = conn.Do("hget", args...)
	return
}

func HMGet(name string, keys []string) (reply []interface{}, err error) {
	conn := GetRedisConn()
	defer conn.Close()
	args := []interface{}{
		name,
	}
	for _, k := range keys {
		args = append(args, k)
	}
	reply, err = redis.Values(conn.Do("hmget", args...))
	return
}

func HMSet(name string, kvs map[string]string) (err error) {
	conn := GetRedisConn()
	defer conn.Close()
	args := []interface{}{
		name,
	}
	for k, v := range kvs {
		args = append(args, k)
		args = append(args, v)
	}
	_, err = conn.Do("hmset", args...)
	return
}

func HKeys(name string) (keys []string, err error) {
	var (
		value []interface{}
	)
	conn := GetRedisConn()
	defer conn.Close()
	keys = make([]string, 0)
	if value, err = redis.Values(conn.Do("hkeys", name)); err == nil {
		for _, v := range value {
			keys = append(keys, string(v.([]byte)))
		}
	}
	return
}

func HVals(name string) (values []string, err error) {
	var (
		value []interface{}
	)
	conn := GetRedisConn()
	defer conn.Close()

	if value, err = redis.Values(conn.Do("hvals", name)); err == nil {
		for _, v := range value {
			values = append(values, string(v.([]byte)))
		}
	}
	return
}

func HDel(name string, field string) (err error) {
	conn := GetRedisConn()
	defer conn.Close()
	_, err = conn.Do("HDEL", name, field)
	return
}

func HGetAll(name string) (kvs map[string]string, err error) {
	var (
		value []interface{}
	)
	conn := GetRedisConn()
	defer conn.Close()
	kvs = make(map[string]string, 0)
	if value, err = redis.Values(conn.Do("hgetall", name)); err == nil {
		for i := 0; i < len(value); i = i + 2 {
			k := string(value[i].([]byte))
			v := string(value[i+1].([]byte))
			kvs[k] = v
		}
	}
	return
}
