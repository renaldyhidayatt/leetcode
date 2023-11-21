package on199

func LevelOrderBottom(root *TreeNode) [][]int {
	tmp := LevelOrder(root)
	res := [][]int{}

	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}

	return res
}
