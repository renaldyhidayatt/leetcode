package on399

type SegmentTree struct {
	arr []int
}

func (st *SegmentTree) Init(nums []int, merge func(int, int) int) {
	size := len(nums)
	st.arr = make([]int, 4*size)
	st.buildTree(nums, 0, 0, size-1, merge)
}

func (st *SegmentTree) buildTree(nums []int, index, left, right int, merge func(int, int) int) {
	if left == right {
		st.arr[index] = nums[left]
		return
	}

	mid := left + (right-left)/2
	st.buildTree(nums, 2*index+1, left, mid, merge)
	st.buildTree(nums, 2*index+2, mid+1, right, merge)
	st.arr[index] = merge(st.arr[2*index+1], st.arr[2*index+2])
}

func (st *SegmentTree) Query(qLeft, qRight int) int {
	return st.queryHelper(0, 0, len(st.arr)/4-1, qLeft, qRight)
}

func (st *SegmentTree) queryHelper(index, sLeft, sRight, qLeft, qRight int) int {
	if qLeft > sRight || qRight < sLeft {
		return 0 // Return a value based on the context of your problem
	}

	if qLeft <= sLeft && qRight >= sRight {
		return st.arr[index]
	}

	mid := sLeft + (sRight-sLeft)/2
	leftVal := st.queryHelper(2*index+1, sLeft, mid, qLeft, qRight)
	rightVal := st.queryHelper(2*index+2, mid+1, sRight, qLeft, qRight)
	return leftVal + rightVal // Modify based on your problem's needs
}

func (st *SegmentTree) Update(i, val int) {
	size := len(st.arr) / 4
	st.updateHelper(0, 0, size-1, i, val)
}

func (st *SegmentTree) updateHelper(index, sLeft, sRight, i, val int) {
	if sLeft == sRight {
		st.arr[index] = val
		return
	}

	mid := sLeft + (sRight-sLeft)/2
	if i <= mid {
		st.updateHelper(2*index+1, sLeft, mid, i, val)
	} else {
		st.updateHelper(2*index+2, mid+1, sRight, i, val)
	}

	st.arr[index] = st.arr[2*index+1] + st.arr[2*index+2] // Modify based on your problem's needs
}

type NumArray struct {
	st SegmentTree
}

func Constructor303(nums []int) NumArray {
	st := SegmentTree{}
	st.Init(nums, func(i, j int) int {
		return i + j
	})
	return NumArray{st: st}
}

func (ma *NumArray) SumRange(i int, j int) int {
	return ma.st.Query(i, j)
}
