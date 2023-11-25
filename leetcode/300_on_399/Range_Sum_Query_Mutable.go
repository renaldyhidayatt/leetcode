package on399

type NumArray307 struct {
	st *SegmentTree
}

func Constructor307(nums []int) NumArray307 {
	st := SegmentTree{}

	st.Init(nums, func(i, j int) int {
		return i + j
	})

	return NumArray307{st: &st}
}

func (this *NumArray307) Update(i int, val int) {
	this.st.Update(i, val)

}

func (this *NumArray307) SumRange(i int, j int) int {
	return this.st.Query(i, j)
}
