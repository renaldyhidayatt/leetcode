package stackgo

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	fmt.Println("Pushed item: ", item)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item
}

func StackMain() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println("popped item:", stack.Pop())
	fmt.Println("stack after popping an element:", stack)
}
