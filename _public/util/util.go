package util

import (
	"reflect"
	"time"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// 获取当前时间字符串
func GetLocalNowTimeStr() string {
	localTime := time.Now().Format("2006-01-02 15:04:05")
	return localTime
}
