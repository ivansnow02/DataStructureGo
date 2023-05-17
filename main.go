package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := []interface{}{1, "2", true, "a", 1000}
	n := 5

	list := newSqList()
	list.createList(arrayToInterface(a), n)
	list.dispList()
	fmt.Println("-------------------------")
	fmt.Println("添加元素1:")
	list.add(1)
	list.dispList()
	fmt.Println("-------------------------")
	fmt.Println("第二个元素是:", list.getElem(2))
	fmt.Println("-------------------------")
	fmt.Println("将第二个元素修改为\"changed\":")
	list.setElem(2, "changed")
	list.dispList()
	fmt.Println("-------------------------")
	fmt.Println("在第3个位置插入\"insert\"")
	list.insert(3, "insert")
	list.dispList()
	fmt.Println("-------------------------")
	fmt.Println("删除第5个元素")
	list.delete(5)
	list.dispList()
	fmt.Println("-------------------------")
	fmt.Println("顺序表长度是:", list.length)
	fmt.Println("-------------------------")
	fmt.Println("a是第几个元素:", list.getNo("a"))
	fmt.Println("-------------------------")
	list.dispList()
}

func arrayToInterface(x interface{}) []interface{} {
	v := reflect.ValueOf(x)          // 获取 x 的反射值
	n := v.Len()                     // 获取 x 的长度
	result := make([]interface{}, n) // 创建一个新的切片
	for i := 0; i < n; i++ {
		result[i] = v.Index(i).Interface() // 获取 x 的每个元素并转换成 interface{}
	}
	return result
}
