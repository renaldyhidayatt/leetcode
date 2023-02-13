package rangequerysum

type NumArray struct {
	sz   int
	nums []int
	sums []int
}

func NewNumArray(nums []int) NumArray {
	sz := len(nums)
	numsArr := make([]int, sz+1)
	sumsArr := make([]int, sz+1)
	na := NumArray{sz: sz, nums: numsArr, sums: sumsArr}
	for i := 0; i < sz; i++ {

	}
	return na
}

func (na *NumArray) update(i int, val int) {
	oldv := na.nums[i+1]
	for idx := i + 1; idx <= na.sz; idx += idx & -idx {
		na.sums[idx] = na.sums[idx] - oldv + val
	}
	na.nums[i+1] = val
}

func (na *NumArray) sumRangee(i int, j int) int {
	return na.sumRange(j+1) - na.sumRange(i)
}

func (na *NumArray) sumRange(i int) int {
	ret := 0
	for idx := i; idx > 0; idx -= idx & -idx {
		ret += na.sums[idx]
	}
	return ret
}
