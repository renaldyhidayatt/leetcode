package main

func sortedArrayToBST(nums []int) *TreeNode {
	return getHelper(nums, 0, len(nums)-1)
}

func getHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = getHelper(nums, start, mid-1)
	node.Right = getHelper(nums, mid+1, end)

	return node
}
