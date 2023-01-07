package main

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0

	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	return 1 + max(ld, rd)
}

func maax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
