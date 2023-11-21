package on199

func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return HasPathSum(root.Left, sum-root.Val) || HasPathSum(root.Right, sum-root.Val)
}
