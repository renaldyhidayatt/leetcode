package linkedlist

import "errors"

type List struct {
	first *Node
	last  *Node
}

type Node struct {
	next  *Node
	prev  *Node
	Value interface{}
}

func NewList(args ...interface{}) *List {
	out := &List{}
	for _, v := range args {
		out.Push(v)
	}

	return out
}

func (l *List) startNode(n *Node) {
	l.first = n
	l.last = n
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	n := Node{Value: v}

	if l.first == nil {
		l.startNode(&n)
		return
	}

	l.first.prev = &n
	n.next = l.first
	l.first = &n
}

func (l *List) Push(v interface{}) {
	n := Node{Value: v}

	if l.last == nil {
		l.startNode(&n)
		return
	}

	n.prev = l.last
	l.last.next = &n
	l.last = &n
}

func (l *List) Shift() (interface{}, error) {
	if l.first == nil {
		return nil, errors.New("cannot shift on empty list")
	}

	if l.first == l.last {
		v := l.first.Value
		l.first, l.last = nil, nil
		return v, nil
	}

	n := l.first

	l.first = n.next
	n.next = nil
	l.first.prev = nil

	return n.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.last == nil {
		return nil, errors.New("cannot pop on empty list")
	}

	if l.first == l.last {
		v := l.last.Value
		l.first, l.last = nil, nil
		return v, nil
	}

	n := l.last

	l.last = n.prev
	n.prev = nil
	l.last.next = nil

	return n.Value, nil
}

func (l *List) Reverse() {
	if l.last == nil {
		return
	}
	l.first, l.last = l.last, l.first

	n := l.first
	for n != nil {
		n.prev, n.next = n.next, n.prev
		n = n.next
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
