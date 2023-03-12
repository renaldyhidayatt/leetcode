package binarytreelevelordertraversalii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	tmp := levelOrder(root)
	res := [][]int{}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelValues)
	}
	return result
}
