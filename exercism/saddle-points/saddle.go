package saddlepoints

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int
type Pair [2]int

func New(s string) (*Matrix, error) {
	rowList := strings.Split(s, "\n")

	m := make(Matrix, 0, len(rowList))

	for _, rowEntry := range rowList {
		row := strings.Fields(rowEntry)

		r := make([]int, 0, len(row))

		for _, element := range row {
			e, err := strconv.Atoi(element)

			if err != nil {
				return &Matrix{}, errors.New("matrix error")
			}
			r = append(r, e)
		}
		m = append(m, r)
	}

	return &m, nil
}

func (m *Matrix) Saddle() []Pair {
	mat := *m
	output := []Pair{}

	iRows, jCols := len(mat), len(mat[0])

	if iRows == 0 || jCols == 0 {
		return output
	}

	booleanMatrix := make(Matrix, iRows)

	for i := 0; i < iRows; i++ {
		booleanRow := make([]int, jCols)
		booleanMatrix[i] = booleanRow
	}

	for i := 0; i < iRows; i++ {
		row := mat[i]
		maxValue := row[0]

		for j := 1; j < jCols; j++ {
			if row[j] > maxValue {
				maxValue = row[j]
			}
		}

		for j := 0; j < jCols; j++ {
			if row[j] == maxValue {
				booleanMatrix[i][j] = 1
			}
		}
	}

	for j := 0; j < jCols; j++ {
		minValue := mat[0][j]

		for i := 1; i < iRows; i++ {
			if mat[i][j] < minValue {
				minValue = mat[i][j]
			}
		}

		for i := 0; i < iRows; i++ {
			if mat[i][j] > minValue {
				booleanMatrix[i][j] = 0
			}
		}
	}

	for i := 0; i < iRows; i++ {
		for j := 0; j < jCols; j++ {
			if booleanMatrix[i][j] == 1 {
				saddle := Pair{i + 1, j + 1}
				output = append(output, saddle)
			}
		}
	}

	return output

}
