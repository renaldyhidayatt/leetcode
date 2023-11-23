package binary

func BitCounter(n uint) int {
	counter := 0

	for n != 0 {
		if n&1 == 1 {
			counter++
		}
		n >>= 1
	}

	return counter
}
