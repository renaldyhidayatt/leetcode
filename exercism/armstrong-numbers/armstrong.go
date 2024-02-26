package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	strNum := strconv.Itoa(n)

	numDigit := len(strNum)

	sum := 0

	for _, digitChar := range strNum {
		digit, _ := strconv.Atoi(string(digitChar))

		sum += int(math.Pow(float64(digit), float64(numDigit)))

	}

	return sum == n
}
