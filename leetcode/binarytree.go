package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr == nil {
			continue
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		stack = append(stack, &TreeNode{Val: curr.Val})
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	for _, node := range stack {
		res = append(res, node.Val)
	}
	return res
}
