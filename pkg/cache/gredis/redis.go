package gredis

import (
	"encoding/json"
	"fmt"
	"github.com/ghjan/learngo/pkg/utils"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool *redis.Pool
)

// newRedisPool : 创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     int(RedisPoolMaxIdle),
		MaxActive:   int(RedisPoolMaxActive),
		IdleTimeout: 300 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			// 1. 打开连接
			c, err := redis.Dial("tcp", RedisHost,
				redis.DialConnectTimeout(time.Duration(RedisDialConnectTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(RedisDialReadTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(RedisDialWriteTimeout)*time.Millisecond))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// 2. 访问认证
			if RedisPass != "" {
				if _, err = c.Do("AUTH", RedisPass); err != nil {
					fmt.Printf("newRedisPool:Dial,AUTH err:%s,config.RedisPass:%s\n",
						err.Error(), RedisPass)
					c.Close()
					return nil, err
				}
			}
			// 3.选择db
			if RedisDB != "" {
				c.Do("SELECT", RedisDB)
			}

			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func InitCachePkg() {
	redisPool = newRedisPool()
}

func GetRedisConn() redis.Conn {
	if redisPool == nil {
		InitCachePkg()
	}
	return redisPool.Get()
}

func Set(key string, data interface{}, time int) (reply interface{}, err error) {
	conn := GetRedisConn()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	reply, err = conn.Do("SET", key, value)
	conn.Do("EXPIRE", key, time)

	return reply, err
}

func Exists(key string) bool {
	conn := GetRedisConn()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := GetRedisConn()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := GetRedisConn()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(patternPart string) (err error) {
	conn := GetRedisConn()
	defer conn.Close()

	if !strings.HasPrefix(patternPart, "*") && !strings.HasSuffix(patternPart, "*") {
		patternPart = "*" + patternPart + "*"
	}
	keys, err := redis.Strings(conn.Do("KEYS", patternPart))
	utils.Checkerr2(err, fmt.Sprintf("LikeDeletes,Do KEYS error,patternPart:%s", patternPart))
	if err == nil {
		for _, key := range keys {
			_, err = Delete(key)
			utils.Checkerr2(err, fmt.Sprintf("LikeDeletes,Delete() error,key:%s", key))
			if err != nil {
				if strings.Contains(err.Error(), "EOF") {
					err = nil
				} else {
					break
				}
			}
		}
	}
	if err != nil && strings.Contains(err.Error(), "EOF") {
		err = nil
	}
	return
}

func DeleteKeys(patternPart string) (msg string, values []string, err error) {
	conn := GetRedisConn()
	defer conn.Close()
	if !strings.HasPrefix(patternPart, "*") && !strings.HasSuffix(patternPart, "*") {
		patternPart = "*" + patternPart + "*"
	}

	values, err = redis.Strings(conn.Do("KEYS", patternPart))
	utils.Checkerr2(err, "DeleteKeys,DO KEYS error")
	if err == nil {
		if values != nil {
			err = conn.Send("MULTI")
			utils.Checkerr2(err, "DeleteKeys,Send MULTI error")
			for i, _ := range values {
				err = conn.Send("DEL", values[i])
				utils.Checkerr2(err, "DeleteKeys,Send DEL error")
			}
			var reply interface{}
			reply, err = conn.Do("EXEC")
			utils.Checkerr2(err, "DeleteKeys,Do EXEC error")
			if reflect.TypeOf(reply).Kind() == reflect.Slice {
				s := reflect.ValueOf(reply)
				for i := 0; i < s.Len(); i++ {
					ele := s.Index(i)
					if reflect.TypeOf(ele).Kind() == reflect.String {
						msg += ele.Interface().(string)
					} else if reflect.TypeOf(ele).Kind() == reflect.Int64 {
						msg += strconv.Itoa(int(ele.Interface().(int64)))
					} else if reflect.TypeOf(ele).Kind() == reflect.Struct {
						msg += fmt.Sprintf("%v", ele)
					} else {
						fmt.Printf("ignored,reflect.TypeOf(reply).Kind():%v,reflect.TypeOf(ele).Kind():%v\n",
							reflect.TypeOf(reply).Kind(), reflect.TypeOf(ele).Kind())
					}
				}
			}
		}
	}
	if err != nil && strings.Contains(err.Error(), "EOF") {
		err = nil
	}
	return
}
