package conversion

import (
	"errors"
	"strconv"
)

func Reverse(str string) string {
	rStr := []rune(str)

	for i, j := 0, len(rStr)-1; i < len(rStr)/2; i, j = i+1, j-1 {
		rStr[i], rStr[j] = rStr[j], rStr[i]
	}
	return string(rStr)
}

func DecimalToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("integer must have +ve value")
	}

	if num == 0 {
		return "0", nil
	}

	var result string = ""

	for num > 0 {
		result += strconv.Itoa(num & 1)
		num >>= 1
	}

	return Reverse(result), nil
}
