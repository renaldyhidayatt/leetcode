package on399

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Rob337(root *TreeNode) int {
	a, b := dfsTreeRob(root)

	return max(a, b)
}

func dfsTreeRob(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}

	l0, l1 := dfsTreeRob(root.Left)
	r0, r1 := dfsTreeRob(root.Right)

	tmp0 := max(l0, l1) + max(r0, r1)

	tmp1 := root.Val + l0 + r0

	return tmp0, tmp1
}
