package main

import "fmt"

const initCap = 5

type SqList struct {
	data             []interface{}
	capacity, length int
}

func newSqList() SqList {
	return SqList{
		data:     make([]interface{}, initCap),
		capacity: initCap,
		length:   0,
	}
}
func (list *SqList) reCap(newCap int) {
	if newCap <= 0 {
		return
	}
	list.capacity = newCap
	var newData = make([]interface{}, newCap)
	for i := 0; i < list.length; i++ {
		newData[i] = list.data[i]
	}
	list.data = newData
}
func (list *SqList) createList(a []interface{}, n int) {
	for i := 0; i < n; i++ {
		if list.length == list.capacity {
			list.reCap(2 * list.length)
		}
		list.data[list.length] = a[i]
		list.length++
	}
}
func (list *SqList) add(e interface{}) {
	if list.length == list.capacity {
		list.reCap(2 * list.length)
	}
	list.data[list.length] = e
	list.length++
}
func (list *SqList) getLength() int {
	return list.length
}
func (list *SqList) getElem(i int) interface{} {
	return list.data[i]
}
func (list *SqList) setElem(i int, e interface{}) bool {
	if i <= 0 || i >= list.length {
		return false
	}
	list.data[i] = e
	return true
}
func (list *SqList) getNo(e interface{}) int {
	for i := 0; i < list.length; i++ {
		if list.data[i] == e {
			return i
		}
	}
	return -1
}
func (list *SqList) insert(i int, e interface{}) bool {
	if i <= 0 || i >= list.length {
		return false
	}
	if list.length == list.capacity {
		list.reCap(2 * list.length)
	}
	for j := list.length; j > i; j-- {
		list.data[j] = list.data[j-1]
	}
	list.data[i] = e
	list.length++
	return true
}
func (list *SqList) delete(i int) bool {
	if i <= 0 || i >= list.length {
		return false
	}
	for j := i; j < list.length-1; j++ {
		list.data[j] = list.data[j+1]
	}
	list.length--
	if list.capacity > initCap && list.length <= list.capacity/4 {
		list.reCap(list.capacity / 2)
	}
	return true
}
func (list *SqList) dispList() {
	fmt.Print("顺序表内容是:[")
	for i := 0; i < list.length-1; i++ {
		fmt.Print(list.data[i], ",")
	}
	fmt.Print(list.data[list.length-1], "]\n")
}
