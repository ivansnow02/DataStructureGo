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
func (l *LinkList) CreateListF(a []interface{}, n int) {
	for i := 0; i < n; i++ {
		l.Insert(0, a[i])
	}
}

// CreateListR 尾插法建表
func (l *LinkList) CreateListR(a []interface{}, n int) {
	for i := 0; i < n; i++ {
		l.Insert(i, a[i])
	}
}

// Insert 在第i个位置插入元素
func (l *LinkList) Insert(i int, e interface{}) bool {
	if i < 0 {
		return false
	}
	s := NewLinkNodeWithData(e)
	p := l.GetI(i - 1)
	if p != nil {
		s.next = p.next
		p.next = s
		if s.next == nil {
			l.tail = s
		}
		l.length++
		return true
	} else {
		return false
	}
}

func (l *LinkList) GetLength() int {
	return l.length
}

// Add 在末尾添加元素
func (l *LinkList) Add(e interface{}) {
	l.Insert(l.length, e)
}

// GetI 返回第i个元素
func (l *LinkList) GetI(i int) *LinkNode {
	p := l.head
	for j := -1; j < i && p != nil; j++ {
		p = p.next
	}
	return p
}

// SetElem 设置序号i的结点值
func (l *LinkList) SetElem(i int, e interface{}) bool {
	if i < 0 {
		return false
	}
	p := l.GetI(i)
	if p != nil {
		p.data = e
		return true
	}
	return false
}

// GetNo 查找第一个为e的元素的序号
func (l *LinkList) GetNo(e interface{}) int {
	no := 0
	p := l.head.next
	for p != nil && p.data != e {
		no++
		p = p.next
	}
	if p == nil {
		return -1
	}
	return no
}

// Delete 在单链表中删除序号i位置的结点
func (l *LinkList) Delete(i int) bool {
	if i < 0 {
		return false
	}
	p := l.GetI(i - 1)
	if p != nil {
		q := p.next
		if q != nil {
			p.next = q.next
			q = nil
			l.length--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// NewLinkList 初始化
func NewLinkList() *LinkList {
	return &LinkList{
		head:   NewLinkNode(),
		tail:   NewLinkNode(),
		length: 0,
	}
}

func (l *LinkList) DispList() {
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
	list.Add("2")
	list.DispList()
	list2 := NewLinkList()
	list2.CreateListR(a, n)
	list2.DispList()
	list2.Add("2")
	list2.DispList()
	list2.Delete(2)
	list2.DispList()
	fmt.Println(list2.GetNo("a"))
	list2.SetElem(3, 123)
	list2.DispList()
}
