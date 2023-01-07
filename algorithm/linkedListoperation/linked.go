package linkedlistoperation

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtBeginning(newData interface{}) {
	newNode := &Node{data: newData}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) InsertAfter(prevNode *Node, newData interface{}) {
	if prevNode == nil {
		fmt.Println("The given previous node must be in the LinkedList")
		return
	}
	newNode := &Node{data: newData}
	newNode.next = prevNode.next
	prevNode.next = newNode
}

func (l *LinkedList) InsertAtEnd(newData interface{}) {
	newNode := &Node{data: newData}
	if l.head == nil {
		l.head = newNode
		return
	}
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

func (l *LinkedList) DeleteNode(position int) {
	if l.head == nil {
		return
	}
	temp := l.head
	if position == 0 {
		l.head = temp.next
		temp = nil
		return
	}
	for i := 0; temp != nil && i < position-1; i++ {
		temp = temp.next
	}
	if temp == nil || temp.next == nil {
		return
	}
	next := temp.next.next
	temp.next = nil
	temp.next = next
}

func (l *LinkedList) Search(key interface{}) bool {
	current := l.head
	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList) SortLinkedList(head *Node) {
	current := head
	index := &Node{}
	if head == nil {
		return
	} else {
		for current != nil {
			index = current.next
			for index != nil {
				if current.data.(int) > index.data.(int) {
					current.data, index.data = index.data, current.data
				}
				index = index.next
			}
			current = current.next
		}
	}
}
