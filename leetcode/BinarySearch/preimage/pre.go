package preimage

func PreImagesSizeFZF(K int) int {
	low, high := 0, 5*K

	for low <= high {
		mid := low + (high-low)>>1
		k := trailingZeroes(mid)
		if k == K {
			return 5
		} else if k > K {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		n /= 5
		count += n
	}
	return count
}
