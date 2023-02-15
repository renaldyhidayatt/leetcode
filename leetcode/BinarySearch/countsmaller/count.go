package countsmaller

import "sort"

type BinaryIndexedTree struct {
	bit []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
	return &BinaryIndexedTree{
		bit: make([]int, n+1),
	}
}

func (b *BinaryIndexedTree) add(i, val int) {
	for i < len(b.bit) {
		b.bit[i] += val
		i += i & -i
	}
}

func (b *BinaryIndexedTree) query(i int) int {
	sum := 0
	for i > 0 {
		sum += b.bit[i]
		i -= i & -i
	}
	return sum
}

func countSmaller1(nums []int) []int {
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

	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bit.query(kth[nums[i]]-1))
		bit.add(kth[nums[i]], 1)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
