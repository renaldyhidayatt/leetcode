package peakindexinamountain

func peakIndexInMountainArray(A []int) int {
	low, high := 0, len(A)-1

	for low <= high {
		mid := low + (high-low)>>1
		if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
			return mid
		}

		if A[mid] > A[mid+1] && A[mid] < A[mid-1] {
			high = mid - 1
		}

		if A[mid] < A[mid+1] && A[mid] > A[mid-1] {
			low = mid + 1
		}
	}

	return 0
}
