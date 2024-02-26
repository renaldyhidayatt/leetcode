package cryptosquare

import (
	"strings"
	"unicode"
)

type Rectangle []string

func Encode(pt string) string {
	pt = strings.ToLower(pt)
	normalizedS := ""
	for _, c := range pt {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			normalizedS += string(c)
		}
	}
	colAmount := CalcuateColumnAmount(len(normalizedS))

	var rectangle Rectangle
	rowStr := ""
	for _, c := range normalizedS {
		if len(rowStr) < colAmount {
			rowStr += string(c)
		}
		if len(rowStr) == colAmount {
			rectangle = append(rectangle, rowStr)
			rowStr = ""
		}
	}
	if len(rowStr) > 0 {
		for len(rowStr) <= colAmount {
			rowStr += " "
		}
		rectangle = append(rectangle, rowStr)
	}

	var resultRectangle Rectangle
	resultRowStr := ""
	colIndex := 0
	for colIndex < colAmount {
		rowIndex := 0
		for rowIndex < len(rectangle) {
			resultRowStr += string(rectangle[rowIndex][colIndex])
			rowIndex++
		}
		resultRectangle = append(resultRectangle, resultRowStr)
		resultRowStr = ""
		colIndex++
	}

	return strings.Join(resultRectangle, " ")
}

func CalcuateColumnAmount(length int) int {
	c := 1
	for c <= length {
		r := c - 1
		for r <= c {
			if r*c >= length {
				return c
			}
			r++
		}
		c++
	}
	return -1
}
