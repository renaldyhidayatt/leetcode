package simplelinkedlist

import "errors"

type list struct {
	values []int
}

func New(values []int) *list {
	return &list{
		values: values,
	}
}

func (l *list) Size() int {
	return len(l.values)
}

func (l *list) Push(element int) {
	l.values = append(l.values, element)
}

func (l *list) Pop() (int, error) {
	if len(l.values) == 0 {
		return 0, errors.New("cant't pop an empty stack")
	}

	last := l.values[len(l.values)-1]

	l.values = l.values[:len(l.values)-1]

	return last, nil
}

func (l *list) Array() []int {
	return l.values
}

func (l *list) Reverse() *list {
	for i, j := 0, len(l.values)-1; i < j; i, j = i+1, j-1 {
		l.values[i], l.values[j] = l.values[j], l.values[i]
	}

	return l
}
