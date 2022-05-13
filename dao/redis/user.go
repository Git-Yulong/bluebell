// Package redis
// @Description:
package redis

import "time"

//
//  StringSet
//  @Description: 设置 string 类型的key
//  @param key
//  @param val
//  @return error
//
func StringSet(key string, val string) error {
	err := rdb.Set(key, val, 300*time.Second).Err()
	return err
}

//
//  GetString
//  @Description:
//  @param key
//  @return string
//  @return error
//
func GetString(key string) (string, error) {
	return rdb.Get(key).Result()
}
