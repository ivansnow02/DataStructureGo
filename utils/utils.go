package utils

import "reflect"

func ArrayToInterface(x interface{}) []interface{} {
	v := reflect.ValueOf(x)          // 获取 x 的反射值
	n := v.Len()                     // 获取 x 的长度
	result := make([]interface{}, n) // 创建一个新的切片
	for i := 0; i < n; i++ {
		result[i] = v.Index(i).Interface() // 获取 x 的每个元素并转换成 interface{}
	}
	return result
}
