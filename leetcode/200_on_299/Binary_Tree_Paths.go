package on299

import "strconv"

func BinaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	tmpLeft := BinaryTreePaths(root.Left)

	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}

	tmpRight := BinaryTreePaths(root.Right)

	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}

	return res
}
