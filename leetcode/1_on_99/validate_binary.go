package on99

import "math"

func IsValidBST(root *TreeNode) bool {
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}
func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}
