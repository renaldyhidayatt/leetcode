package stringtointeger

import (
	"fmt"
	"math"
	"unicode"
)

const maxDiv1 = math.MaxInt32 / 10

func MyAtoi(str string) int {
	i, n := 0, len(str)

	for i < n && unicode.IsSpace(rune(str[i])) {
		i++
	}

	sign := 1

	if i < n && str[i] == '+' {
		i++
	} else if i < n && str[i] == '-' {
		sign = -1
		i++
	}

	num := 0

	for i < n && unicode.IsDigit(rune(str[i])) {
		digit := int(str[i] - '0')
		if num > maxDiv1 || num == maxDiv1 && digit >= 8 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		num = num*10 + digit
		i++
	}

	return sign * num
}

func RunningMyAtoi() {
	fmt.Println(MyAtoi("42"))
}
