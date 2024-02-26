package sort

import "algoritmAndDs/constraints"

const runSizeThreshold = 8

func Timsort[T constraints.Ordered](data []T) []T {
	runSize := calculateRunSize(len(data))
	insertionSortRuns(data, runSize)
	mergeRuns(data, runSize)
	return data
}

func calculateRunSize(dataLength int) int {
	remainder := 0
	for dataLength >= runSizeThreshold {
		if dataLength%2 == 1 {
			remainder = 1
		}

		dataLength = dataLength / 2
	}

	return dataLength + remainder
}

func insertionSortRuns[T constraints.Ordered](data []T, runSize int) {
	for lower := 0; lower < len(data); lower += runSize {
		upper := lower + runSize
		if upper >= len(data) {
			upper = len(data)
		}

		Insertion(data[lower:upper])
	}
}

func mergeRuns[T constraints.Ordered](data []T, runSize int) {
	for size := runSize; size < len(data); size *= 2 {
		for lowerBound := 0; lowerBound < len(data); lowerBound += size * 2 {
			middleBound := lowerBound + size - 1
			upperBound := lowerBound + 2*size - 1
			if len(data)-1 < upperBound {
				upperBound = len(data) - 1
			}

			mergeRun(data, lowerBound, middleBound, upperBound)
		}
	}
}

func mergeRun[T constraints.Ordered](data []T, lower, mid, upper int) {
	left := data[lower : mid+1]
	right := data[mid+1 : upper+1]
	merged := merge(left, right)
	for i, value := range merged {
		data[lower+i] = value
	}
}
