package main

type Node struct {
	item interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(item interface{}) {
	newNode := &Node{item: item}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
}

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		println(currentNode.item)
		currentNode = currentNode.next
	}
}

func LinkedListMain() {
	linkedList := &LinkedList{}
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Print()
}
