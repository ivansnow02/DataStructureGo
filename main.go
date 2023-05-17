package main

import (
	"DataStructure/List"
	"fmt"
	"reflect"
)

func main() {

	a := []interface{}{1, "2", true, "a", 1000}
	n := 5

	list := List.NewSqList()
	list.CreateList(arrayToInterface(a), n)
	list.DispList()
	fmt.Println("-------------------------")
	fmt.Println("添加元素1:")
	list.Add(1)
	list.DispList()
	fmt.Println("-------------------------")
	fmt.Println("第二个元素是:", list.GetElem(2))
	fmt.Println("-------------------------")
	fmt.Println("将第二个元素修改为\"changed\":")
	list.SetElem(2, "changed")
	list.DispList()
	fmt.Println("-------------------------")
	fmt.Println("在第3个位置插入\"insert\"")
	list.Insert(3, "insert")
	list.DispList()
	fmt.Println("-------------------------")
	fmt.Println("删除第5个元素")
	list.Delete(5)
	list.DispList()
	fmt.Println("-------------------------")
	fmt.Println("顺序表长度是:", list.Length)
	fmt.Println("-------------------------")
	fmt.Println("a是第几个元素:", list.GetNo("a"))
	fmt.Println("-------------------------")
	list.DispList()
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
