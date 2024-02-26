package dynamic

func LongestIncreasingSubsequenceGreedy(nums []int) int {
	longestIncreasingSubsequence := make([]int, 0)

	for _, num := range nums {
		leftmostIndex := lowerBound(longestIncreasingSubsequence, num)

		if leftmostIndex == len(longestIncreasingSubsequence) {
			longestIncreasingSubsequence = append(longestIncreasingSubsequence, num)
		} else {
			longestIncreasingSubsequence[leftmostIndex] = num
		}
	}

	return len(longestIncreasingSubsequence)
}

func lowerBound(arr []int, val int) int {
	searchWindowLeft, searchWindowRight := 0, len(arr)-1

	for searchWindowLeft <= searchWindowRight {
		middle := (searchWindowLeft + searchWindowRight) / 2

		if arr[middle] < val {
			searchWindowLeft = middle + 1
		} else {
			searchWindowRight = middle - 1
		}
	}

	return searchWindowRight + 1
}
