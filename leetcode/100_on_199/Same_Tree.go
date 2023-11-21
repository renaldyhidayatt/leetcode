package on199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
