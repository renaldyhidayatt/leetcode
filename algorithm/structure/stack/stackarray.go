package stack

type Array[T any] struct {
	elements []T
}

func NewStack[T any]() *Array[T] {
	return &Array[T]{}
}

func (s *Array[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Array[T]) Length() int {
	return len(s.elements)
}

func (s *Array[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T

		return zeroValue
	}

	return s.elements[len(s.elements)-1]
}

func (s *Array[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Array[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T

		return zeroValue
	}

	index := len(s.elements) - 1
	popped := s.elements[index]

	s.elements = s.elements[:index]

	return popped
}
