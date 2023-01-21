package palingdromenumber

import "fmt"

func IsPalingdrom(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}

func RunningIsPalingDrom() {
	fmt.Println(IsPalingdrom(121))
}
