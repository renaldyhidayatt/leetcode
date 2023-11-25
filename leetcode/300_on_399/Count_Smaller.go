package on399

import "sort"

type BinaryIndexedTree struct {
	tree []int
}

func (bit *BinaryIndexedTree) Init(size int) {
	bit.tree = make([]int, size+1)
}

func (bit *BinaryIndexedTree) Query(index int) int {
	index++
	result := 0
	for index > 0 {
		result += bit.tree[index]
		index -= index & -index
	}
	return result
}

func (bit *BinaryIndexedTree) Add(index, val int) {
	index++
	for index < len(bit.tree) {
		bit.tree[index] += val
		index += index & -index
	}
}

func CountSmaller(nums []int) []int {
	allNums, res := make([]int, len(nums)), []int{}

	copy(allNums, nums)

	sort.Ints(allNums)

	k := 1

	kth := map[int]int{allNums[0]: k}

	for i := 1; i < len(allNums); i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	bit := BinaryIndexedTree{}

	bit.Init(k)

	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bit.Query(kth[nums[i]]-1))
		bit.Add(kth[nums[i]], 1)
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
