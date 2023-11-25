package on399

import "math/rand"

type Solution384 struct {
	nums []int
}

func Constructor384(nums []int) Solution384 {
	return Solution384{
		nums: nums,
	}
}

func (this *Solution384) Reset() []int {
	return this.nums
}

func (this *Solution384) Shuffle() []int {
	arr := make([]int, len(this.nums))
	copy(arr, this.nums)

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}
