package maxlenofrepeatedsub

const primeRK = 16777619

func findLength(A []int, B []int) int {
	low, high := 0, min(len(A), len(B))
	for low < high {
		mid := (low + high + 1) >> 1
		if hasRepeated(A, B, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func hashSlice(arr []int, length int) []int {
	hash, pl, h := make([]int, len(arr)-length+1), 1, 0

	for i := 0; i < length-1; i++ {
		pl *= primeRK
	}

	for i, v := range arr {
		h = h*primeRK + v
		if i >= length-1 {
			hash[i-length+1] = h
			h -= pl * arr[i-length+1]
		}

	}
	return hash
}

func hasSamePrefix(A, B []int, length int) bool {
	for i := 0; i < length; i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}

func hasRepeated(A, B []int, length int) bool {
	hs := hashSlice(A, length)
	hashToOffset := make(map[int][]int, len(hs))
	for i, h := range hs {
		hashToOffset[h] = append(hashToOffset[h], i)
	}
	for i, h := range hashSlice(B, length) {
		if offsets, ok := hashToOffset[h]; ok {
			for _, offset := range offsets {
				if hasSamePrefix(A[offset:], B[i:], length) {
					return true
				}
			}
		}
	}
	return false
}
