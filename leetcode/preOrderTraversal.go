package main

type treenode struct {
	Val   int
	Left  *treenode
	Right *treenode
}

func preorderTravesal(root *treenode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*treenode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	return res
}
