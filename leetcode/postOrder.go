package main

func postorderTraversal(root *TreeNode) []int {
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
		res = append(res, curr.Val)
		stack = append(stack, nil)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	reverse(res)

	return res
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
