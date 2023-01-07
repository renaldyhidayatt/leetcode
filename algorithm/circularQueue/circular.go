package circularqueue

import "fmt"

type MyCircularQueue struct {
	k     int
	queue []int
	head  int
	tail  int
}

func (cq *MyCircularQueue) Enqueue(data int) {
	if (cq.tail+1)%cq.k == cq.head {
		fmt.Println("The circular queue is full")
		return
	}

	if cq.head == -1 {
		cq.head = 0
		cq.tail = 0
		cq.queue[cq.tail] = data
	} else {
		cq.tail = (cq.tail + 1) % cq.k
		cq.queue[cq.tail] = data
	}
}

func (cq *MyCircularQueue) Dequeue() int {
	if cq.head == -1 {
		fmt.Println("The circular queue is empty")
		return 0
	}

	if cq.head == cq.tail {
		temp := cq.queue[cq.head]
		cq.head = -1
		cq.tail = -1
		return temp
	}

	temp := cq.queue[cq.head]
	cq.head = (cq.head + 1) % cq.k
	return temp
}

func (cq *MyCircularQueue) PrintCQueue() {
	if cq.head == -1 {
		fmt.Println("No element in the circular queue")
		return
	}

	if cq.tail >= cq.head {
		for i := cq.head; i <= cq.tail; i++ {
			fmt.Print(cq.queue[i], " ")
		}
		for i := 0; i <= cq.tail; i++ {
			fmt.Print(cq.queue[i], " ")
		}
		fmt.Println()
	}
}

func CircularQueue() {
	cq := MyCircularQueue{5, make([]int, 5), -1, -1}
	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)
	cq.Enqueue(4)
	cq.Enqueue(5)
	fmt.Println("Initial queue")
	cq.PrintCQueue()

	cq.Dequeue()
	fmt.Println("After removing an element from the queue")
	cq.PrintCQueue()
}
