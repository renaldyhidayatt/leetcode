package conversion

import (
	"errors"
	"strings"
)

type numeral struct {
	val int
	sym string
}

var nums = []numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func RomanToInt(input string) (int, error) {
	if input == "" {
		return 0, nil
	}
	var output int

	for _, n := range nums {
		for strings.HasPrefix(input, n.sym) {
			output += n.val
			input = input[len(n.sym):]
		}
	}

	if len(input) > 0 {
		return 0, errors.New("invalid roman numeral")
	}

	return output, nil
}
