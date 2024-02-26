package prime

func Factorize(n int64) map[int64]int64 {
	result := make(map[int64]int64)

	for i := int64(2); i*i <= n; i++ {
		for {
			if n%i != 0 {
				break
			}
			result[i] += 1
			n /= i
		}
	}

	if n > 1 {
		result[n] += 1
	}

	return result
}
