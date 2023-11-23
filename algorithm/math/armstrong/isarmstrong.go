package armstrong

import (
	"math"
	"strconv"
)

func IsArmstrong(number int) bool {
	var rightMost int

	var sum int = 0
	var tempNum int = number

	lenght := float64(len(strconv.Itoa(number)))

	for tempNum > 0 {
		rightMost = tempNum % 10
		sum += int(math.Pow(float64(rightMost), lenght))

		tempNum /= 10
	}

	return number == sum
}
