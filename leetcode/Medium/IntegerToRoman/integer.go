package integertoroman

import (
	"bytes"
	"fmt"
)

func intToRoman(num int) string {
	mapping := map[int]string{
		1: "I", 5: "V", 10: "X",
		50: "L", 100: "C", 500: "D", 1000: "M",
		4: "IV", 9: "IX", 40: "XL", 90: "XC",
		400: "CD", 900: "CM",
	}

	sequance := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var buf bytes.Buffer
	for _, base := range sequance {
		for num >= base {
			buf.WriteString(mapping[base])
			num -= base
		}
	}

	return buf.String()
}

func RunningIntToRoman() {
	fmt.Println(intToRoman(90))
}
