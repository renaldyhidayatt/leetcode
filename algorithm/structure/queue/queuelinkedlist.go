package queue

type Node struct {
	Data any
	Next *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *Queue) Enqueue(n any) {
	var newNode Node

	newNode.Data = n

	if ll.tail != nil {
		ll.tail.Next = &newNode
	}

	ll.tail = &newNode

	if ll.head == nil {
		ll.head = &newNode
	}

	ll.length++
}

func (ll *Queue) Dequeue() any {
	if ll.IsEmpty() {
		return -1
	}

	data := ll.head.Data

	ll.head = ll.head.Next

	if ll.head == nil {
		ll.tail = nil
	}

	ll.length--

	return data
}

func (ll *Queue) IsEmpty() bool {
	return ll.length == 0
}

func (ll *Queue) Len() int {
	return ll.length
}

func (ll *Queue) FrontQueue() any {
	return ll.head.Data
}

func (ll *Queue) BackQueue() any {
	return ll.tail.Data
}
