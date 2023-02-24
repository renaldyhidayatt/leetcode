package superegg

func superEggDrop(K int, N int) int {
	low, high := 1, N
	for low < high {
		mid := low + (high-low)>>1
		if counterF(K, N, mid) >= N {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func counterF(k, n, mid int) int {
	res, sum := 1, 0
	for i := 1; i <= k && sum < n; i++ {
		res *= mid - i + 1
		res /= i
		sum += res
	}

	return sum
}
