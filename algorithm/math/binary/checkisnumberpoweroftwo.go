package binary

func IsPowerOfTwo(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

func IsPowerOfTwoLeftShift(number uint) bool {
	for p := uint(1); p <= number; p = p << 1 {
		if number == p {
			return true
		}
	}
	return false
}
