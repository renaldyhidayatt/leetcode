package on199

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	for pos, node := range inorder {
		if node == root.Val {
			root.Left = BuildTree(preorder[1:pos+1], inorder[:pos])
			root.Right = BuildTree(preorder[pos+1:], inorder[pos+1:])
		}
	}

	return root
}
