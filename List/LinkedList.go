package List

import "fmt"

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
	head *LinkNode
	tail *LinkNode
}

func (l LinkList) CreateListF(a []interface{}, n int) {
	for i := 0; i < n; i++ {
		s := NewLinkNodeWithData(a[i])
		s.next = l.head.next
		l.head.next = s
	}

}

func NewLinkList() *LinkList {
	return &LinkList{
		head: NewLinkNode(),
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
}
