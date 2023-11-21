package on199

func preorder(root *TreeNode, list *[]int) {
	if root != nil {
		*list = append(*list, root.Val)
		preorder(root.Left, list)
		preorder(root.Right, list)
	}
}

func Flatten(root *TreeNode) {
	list, cur := []int{}, &TreeNode{}

	preorder(root, &list)

	cur = root

	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{Val: list[i], Left: nil, Right: nil}
		cur = cur.Right
	}

	return
}
