package main

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
	return isValidBSTHelper(root.Left, minVal, root.Val) && isValidBSTHelper(root.Right, root.Val, maxVal)
}
