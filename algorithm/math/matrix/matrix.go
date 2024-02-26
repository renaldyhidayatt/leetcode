package matrix

import (
	"algoritmAndDs/constraints"
	"errors"
	"sync"
)

type Matrix[T constraints.Integer] struct {
	elements [][]T
	rows     int
	columns  int
}

func New[T constraints.Integer](rows, columns int, initial T) Matrix[T] {
	if rows < 0 || columns < 0 {
		return Matrix[T]{}
	}

	elements := make([][]T, rows)

	var wg sync.WaitGroup

	wg.Add(rows)

	for i := range elements {
		go func(i int) {
			defer wg.Done()

			elements[i] = make([]T, columns)

			for j := range elements[i] {
				elements[i][j] = initial
			}
		}(i)
	}

	wg.Wait()

	return Matrix[T]{elements: elements, rows: rows, columns: columns}
}

func NewFromElements[T constraints.Integer](elements [][]T) (Matrix[T], error) {
	if !IsValid[T](elements) {
		return Matrix[T]{}, errors.New("rows have different numbers of columns")
	}

	rows := len(elements)

	if rows == 0 {
		return Matrix[T]{}, nil
	}

	columns := len(elements[0])

	matrix := Matrix[T]{
		elements: make([][]T, rows),
		rows:     rows,
		columns:  columns,
	}

	for i := range matrix.elements {
		matrix.elements[i] = make([]T, columns)
		copy(matrix.elements[i], elements[i])
	}

	return matrix, nil
}

func (m Matrix[T]) Get(row, col int) (T, error) {
	if row < 0 || row >= m.rows || col < 0 || col >= m.columns {
		var zeroVal T
		return zeroVal, errors.New("index out of range")
	}

	return m.elements[row][col], nil
}

func (m Matrix[T]) Set(row, col int, val T) error {
	if row < 0 || row >= m.rows || col < 0 || col >= m.columns {
		return errors.New("index out of range")
	}

	m.elements[row][col] = val

	return nil
}

func (m Matrix[T]) Rows() int {
	return len(m.elements)
}

func (m Matrix[T]) Columns() int {
	if len(m.elements) == 0 {
		return 0
	}

	return len(m.elements[0])
}
