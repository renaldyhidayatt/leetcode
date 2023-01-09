package perfectbinarytree

import "fmt"

type newNode struct {
	key   int
	left  *newNode
	right *newNode
}

func calculateDepth(node *newNode) int {
	d := 0
	for node != nil {
		d++
		node = node.left
	}
	return d
}

func isPerfect(root *newNode, d int, level int) bool {
	if root == nil {
		return true
	}
	if root.left == nil && root.right == nil {
		return d == level+1
	}
	if root.left == nil || root.right == nil {
		return false
	}
	return isPerfect(root.left, d, level+1) && isPerfect(root.right, d, level+1)
}

func PerfectBinaryMain() {

	root := &newNode{key: 1}
	root.left = &newNode{key: 2}
	root.right = &newNode{key: 3}
	root.left.left = &newNode{key: 4}
	root.left.right = &newNode{key: 5}

	if isPerfect(root, calculateDepth(root), 0) {
		fmt.Println("The tree is a perfect binary tree")
	} else {
		fmt.Println("The tree is not a perfect binary tree")
	}

}
