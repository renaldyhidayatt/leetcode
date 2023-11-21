package on199

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
