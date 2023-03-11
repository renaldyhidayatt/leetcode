package minnumberofkconsecutive

func minKBitFlips(A []int, K int) int {
	flippedTime, count := 0, 0

	for i := 0; i < len(A); i++ {
		if i >= K && A[i-K] == 2 {
			flippedTime--
		}

		if flippedTime%2 == A[i] {
			if i+K > len(A) {
				return -1
			}
			A[i] = 2
			flippedTime++
			count++
		}
	}

	return count
}
