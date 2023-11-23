package on199

func PreorderTraversal(root *TreeNode) []int {
	var result []int
	preorderBinary(root, &result)
	return result
}

func preorderBinary(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}
