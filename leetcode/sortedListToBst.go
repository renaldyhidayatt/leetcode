package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBst(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	size := 0
	pos := head

	for pos != nil {
		pos = pos.Next
		size++
	}

	return inorderHelper(head, 0, size-1)
}

func inorderHelper(node *ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	left := inorderHelper(node, start, mid-1)
	root := &TreeNode{Val: node.Val}

	root.Left = left
	node = node.Next
	root.Right = inorderHelper(node, mid+1, end)
	return root
}
