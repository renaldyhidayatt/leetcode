package binarytreeorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := make([][]int, 0)

	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, 1)

		for i := 0; i < 1; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, tmp)
	}

	return res
}
