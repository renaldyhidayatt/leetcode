package on199

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}

	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}

	return min(MinDepth(root.Left), MinDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
