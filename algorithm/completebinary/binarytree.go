package completebinary

import "fmt"

type Node struct {
	item  int
	left  *Node
	right *Node
}

func countNodes(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.left) + countNodes(root.right)
}

func isComplete(root *Node, index int, numberNodes int) bool {
	if root == nil {
		return true
	}
	if index >= numberNodes {
		return false
	}
	return isComplete(root.left, 2*index+1, numberNodes) && isComplete(root.right, 2*index+2, numberNodes)
}

func BinaryTreeCompleteMain() {
	root := &Node{item: 1}
	root.left = &Node{item: 2}
	root.right = &Node{item: 3}
	root.left.left = &Node{item: 4}
	root.left.right = &Node{item: 5}
	root.right.left = &Node{item: 6}

	nodeCount := countNodes(root)
	index := 0

	if isComplete(root, index, nodeCount) {
		fmt.Println("The tree is a complete binary tree")
	} else {
		fmt.Println("The tree is not a complete binary tree")
	}

}
