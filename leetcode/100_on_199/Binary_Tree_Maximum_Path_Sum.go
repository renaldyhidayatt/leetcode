package on199

import "math"

func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := math.MinInt32

	GetPathSum(root, &max)

	return max
}

func GetPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}

	left := GetPathSum(root.Left, maxSum)
	right := GetPathSum(root.Right, maxSum)

	currMax := max(max(left+root.Val, right+root.Val), root.Val)
	*maxSum = max(*maxSum, max(currMax, left+right+root.Val))

	return currMax
}
