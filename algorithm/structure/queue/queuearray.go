package queue

var ListQueue []any

func Enqueue(value any) {
	ListQueue = append(ListQueue, value)
}

func Dequeue() any {
	value := ListQueue[0]
	ListQueue = ListQueue[1:]
	return value
}

func FrontQueue() any {
	return ListQueue[0]
}

func BackQueue() any {
	return ListQueue[len(ListQueue)-1]
}

func LenQueue() int {
	return len(ListQueue)
}

func IsEmptyQueue() bool {
	return len(ListQueue) == 0
}
