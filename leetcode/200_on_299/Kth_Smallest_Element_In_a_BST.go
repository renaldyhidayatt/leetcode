package on299

func KthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0

	Inorder230(root, k, &count, &res)

	return res
}

func Inorder230(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		Inorder230(node.Left, k, count, ans)

		*count++
		if *count == k {
			*ans = node.Val

			return
		}

		Inorder230(node.Right, k, count, ans)
	}
}
