package linklist

import "fmt"

type Node struct {
	Next  *Node
	Value interface{}
}

type LinkedList struct {
	Head     *Node
	capacity int
	length   int
}

func NewLinkedList(capacity int) *LinkedList {
	return &LinkedList{
		capacity: capacity,
	}
}

func (l *LinkedList) Insert(val interface{}) error {
	if l.length == l.capacity {
		return fmt.Errorf("list is full")
	}
	if l.Head == nil {
		l.Head = &Node{Value: val}
	} else {
		cur := l.Head
		for i := 0; i < l.length-1; i++ {
			cur = cur.Next
		}
		cur.Next = &Node{Value: val}
	}
	l.length++
	return nil
}

func (L *LinkedList) Remove(val interface{}) error {
	return nil
}

func (L *LinkedList) GetLen() int {
	return L.length
}
