package queue

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Removing an element from the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return 0
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Display the queue
func (q *Queue) Display() {
	fmt.Println(q.items)
}

// Get the size of the queue
func (q *Queue) Size() int {
	return len(q.items)
}

func QueueMain() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	q.Display()

	q.Dequeue()

	fmt.Println("After removing an element")
	q.Display()
}
