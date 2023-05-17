package List

import "fmt"

const initCap = 5

type SqList struct {
	data             []interface{}
	capacity, Length int
}

func NewSqList() SqList {
	return SqList{
		data:     make([]interface{}, initCap),
		capacity: initCap,
		Length:   0,
	}
}
func (list *SqList) reCap(newCap int) {
	if newCap <= 0 {
		return
	}
	list.capacity = newCap
	var newData = make([]interface{}, newCap)
	for i := 0; i < list.Length; i++ {
		newData[i] = list.data[i]
	}
	list.data = newData
}
func (list *SqList) CreateList(a []interface{}, n int) {
	for i := 0; i < n; i++ {
		if list.Length == list.capacity {
			list.reCap(2 * list.Length)
		}
		list.data[list.Length] = a[i]
		list.Length++
	}
}
func (list *SqList) Add(e interface{}) {
	if list.Length == list.capacity {
		list.reCap(2 * list.Length)
	}
	list.data[list.Length] = e
	list.Length++
}
func (list *SqList) getLength() int {
	return list.Length
}
func (list *SqList) GetElem(i int) interface{} {
	return list.data[i]
}
func (list *SqList) SetElem(i int, e interface{}) bool {
	if i <= 0 || i >= list.Length {
		return false
	}
	list.data[i] = e
	return true
}
func (list *SqList) GetNo(e interface{}) int {
	for i := 0; i < list.Length; i++ {
		if list.data[i] == e {
			return i
		}
	}
	return -1
}
func (list *SqList) Insert(i int, e interface{}) bool {
	if i <= 0 || i >= list.Length {
		return false
	}
	if list.Length == list.capacity {
		list.reCap(2 * list.Length)
	}
	for j := list.Length; j > i; j-- {
		list.data[j] = list.data[j-1]
	}
	list.data[i] = e
	list.Length++
	return true
}
func (list *SqList) Delete(i int) bool {
	if i <= 0 || i >= list.Length {
		return false
	}
	for j := i; j < list.Length-1; j++ {
		list.data[j] = list.data[j+1]
	}
	list.Length--
	if list.capacity > initCap && list.Length <= list.capacity/4 {
		list.reCap(list.capacity / 2)
	}
	return true
}
func (list *SqList) DispList() {
	fmt.Print("顺序表内容是:[")
	for i := 0; i < list.Length-1; i++ {
		fmt.Print(list.data[i], ",")
	}
	fmt.Print(list.data[list.Length-1], "]\n")
}
