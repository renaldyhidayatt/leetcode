package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (m Matrix, err error) {

	lines := strings.Split(s, "\n")

	m = make([][]int, len(lines))

	var rowLength int

	for i, line := range lines {

		unparsed := strings.Fields(line)

		if i == 0 {

			rowLength = len(unparsed)

		}

		if len(unparsed) != rowLength {
			return nil, errors.New("")

		}

		for _, n := range unparsed {

			parsed, convErr := strconv.Atoi(n)

			if convErr != nil {

				return nil, convErr

			}

			m[i] = append(m[i], parsed)

		}

	}

	return

}

// Cols and Rows must return the results without affecting the matrix.

func (m Matrix) Cols() (cols [][]int) {

	cols = make([][]int, len(m[0]))

	for i := range m[0] {

		for _, r := range m {

			cols[i] = append(cols[i], r[i])

		}

	}

	return

}

func (m Matrix) Rows() (rows [][]int) {

	rows = make([][]int, len(m))

	for i := range rows {

		rows[i] = append(rows[i], m[i]...)

	}

	return

}

func (m Matrix) Set(row, col, val int) bool {

	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {

		return false

	}

	m[row][col] = val

	return true

}
