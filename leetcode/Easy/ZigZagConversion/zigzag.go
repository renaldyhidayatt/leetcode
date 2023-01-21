package zigzagconversion

import "fmt"

func Convert(s string, nRows int) string {
	r := make([]string, nRows)
	row := 0
	step := 1
	for i := 0; i < len(s); i++ {
		if row == nRows-1 {
			step = -1
		}
		if row == 0 {
			step = 1
		}
		r[row] += string(s[i])
		row += step
	}
	result := ""
	for i := 0; i < nRows; i++ {
		result += r[i]
	}
	return result
}

func RunningZigzagConversion() {
	fmt.Println(Convert("PAYPALISHIRING", 3))
}
