package on199

func PostorderTraversal(root *TreeNode) []int {
	var result []int
	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}
