package link

import (
	"errors"
	"sync"
)

type Node struct {
	ID  int
	Pop bool // 判断是否被pop过

	Data interface{}

	Next *Node
}

type LinkQueue struct {
	root *Node

	head *Node
	tail *Node

	size int
	len  int

	full bool // 判断链表长度是否达到指定深度

	mu sync.Mutex
}

func NewLinkQ(s int) *LinkQueue {
	q := &LinkQueue{size: s, root: &Node{}}
	return q
}

func (l *LinkQueue) Push(data interface{}) error {

	if l.size < 1 {
		return errors.New("[link] queue size is 0")
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// 替换数据，不在创建新节点
	if l.full {
		end := l.tail.Next
		if end.Pop == false {
			return errors.New("[link] queue data is full")
		}

		end.Data = data
		end.Pop = false
		l.tail = end

		l.len++
		return nil
	}

	elem := new(Node)
	elem.Data = data
	if l.root.Next == nil {
		l.root.Next = elem
	}

	end := l.tail
	if end != nil {
		elem.ID = end.ID + 1
		end.Next = elem
	}
	l.tail = elem
	l.tail.Next = l.root.Next
	if l.tail.ID == l.size-1 {
		l.full = true
	}

	l.len++
	return nil
}

func (l *LinkQueue) Pop() (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.root.Next == nil {
		return nil, false
	}
	if l.head == nil {
		l.head = l.root.Next
	}
	elem := l.head
	if elem.Pop == true {
		return nil, false
	}

	elem.Pop = true
	l.head = elem.Next
	l.len--
	return elem.Data, true

}

func (l *LinkQueue) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.len
}

func (l *LinkQueue) IsEmpty() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.len == 0
}
