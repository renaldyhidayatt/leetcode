package kthsmallestprimefraction

func KthSmallestPrimeFraction(A []int, K int) []int {
	low, high, n := 0.0, 1.0, len(A)

	for {
		mid, count, p, q, j := (high+low)/2.0, 0, 0, 1, 0
		for i := 0; i < n; i++ {
			for j < n && float64(A[i]) > float64(mid)*float64(A[j]) {
				j++
			}
			count += n - j
			if j < n && q*A[i] > p*A[j] {
				p = A[i]
				q = A[j]
			}
		}
		if count == K {
			return []int{p, q}
		} else if count < K {
			low = mid
		} else {
			high = mid
		}
	}
}
