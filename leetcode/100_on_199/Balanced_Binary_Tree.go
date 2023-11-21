package on199

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHight := depth(root.Left)
	rightHight := depth(root.Right)

	return abs(leftHight-rightHight) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
