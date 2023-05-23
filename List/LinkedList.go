package List

import (
	"fmt"
)

type LinkNode struct {
	data interface{}
	next *LinkNode
}

func NewLinkNode() *LinkNode {
	return &LinkNode{
		next: nil,
	}
}
func NewLinkNodeWithData(d interface{}) *LinkNode {
	return &LinkNode{
		data: d,
		next: nil,
	}
}

type LinkList struct {
	head   *LinkNode
	tail   *LinkNode
	length int
}

// CreateListF 头插法建表
func (l LinkList) CreateListF(a []interface{}, n int) {
	newNode := NewLinkNodeWithData(a[0])
	l.head.next = newNode
	l.tail = newNode
	for i := 1; i < n; i++ {
		l.Insert(0, a[i])
	}
}

// CreateListR 尾插法建表
func (l LinkList) CreateListR(a []interface{}, n int) {
	newNode := NewLinkNodeWithData(a[0])
	l.head.next = newNode
	l.tail = newNode
	for i := 1; i < n; i++ {
		newNode := NewLinkNodeWithData(a[i])
		l.tail.next = newNode
		l.tail = newNode
	}
	l.tail.next = nil
}

// Insert 在第i个位置插入元素
func (l LinkList) Insert(i int, e interface{}) bool {
	if i < 0 {
		return false
	}
	s := NewLinkNodeWithData(e)
	p := l.GetI(i - 1)
	if p != nil {
		s.next = p.next
		p.next = s
		l.length++
		return true
	} else {
		return false
	}
}

// GetI 返回第i个元素
func (l LinkList) GetI(i int) *LinkNode {
	p := l.head
	for j := -1; j < i && p != nil; j++ {
		p = p.next
	}
	return p
}

// NewLinkList 初始化
func NewLinkList() *LinkList {
	return &LinkList{
		head:   NewLinkNode(),
		tail:   NewLinkNode(),
		length: 0,
	}
}

func (l LinkList) DispList() {
	p := l.head.next
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}
	fmt.Println()
}

func TestLinkedList() {
	a := []interface{}{1, "2", true, "a", 1000}
	n := 5
	list := NewLinkList()
	list.CreateListF(a, n)
	list.DispList()
	list2 := NewLinkList()
	list2.CreateListR(a, n)
	list2.DispList()
	//fmt.Println(list.GetI(1))
}
