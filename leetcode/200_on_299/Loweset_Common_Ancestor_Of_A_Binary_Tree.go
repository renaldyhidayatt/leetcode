package on299

func LowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == q {
		return root
	}

	left := LowestCommonAncestor236(root.Left, p, q)
	right := LowestCommonAncestor236(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
