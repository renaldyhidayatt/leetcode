package on499

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + SumOfLeftLeaves(root.Right)
	}

	return SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)
}
