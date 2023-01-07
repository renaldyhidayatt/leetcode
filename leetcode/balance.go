package main

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if getDepth(root) < 1 {
		return false
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 1
	}

	ld := getDepth(node.Left)

	if ld < 0 {
		return -1

	}
	rd := getDepth(node.Right)

	if abs(ld-rd) > 1 {
		return -1
	}

	return maax(ld, rd) + 1
}

func maax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
