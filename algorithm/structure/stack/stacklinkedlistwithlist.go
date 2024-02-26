package stack

import (
	"container/list"
	"fmt"
)

type SList struct {
	Stack *list.List
}

func (sl *SList) Push(val any) {
	sl.Stack.PushFront(val)
}

func (sl *SList) Peek() (any, error) {
	if !sl.IsEmpty() {
		element := sl.Stack.Front()

		return element.Value, nil
	}

	return "", fmt.Errorf("stack list is empty")
}

func (sl *SList) Pop() (any, error) {
	if !sl.IsEmpty() {
		element := sl.Stack.Front()

		sl.Stack.Remove(element)

		return element.Value, nil
	}

	return "", fmt.Errorf("stack list is empty")
}

func (sl *SList) IsEmpty() bool {
	return sl.Stack.Len() == 0
}
