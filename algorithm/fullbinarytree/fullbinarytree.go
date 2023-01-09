package fullbinarytree

import "fmt"

type Node struct {
	item       int
	leftChild  *Node
	rightChild *Node
}

func isFullTree(root *Node) bool {
	if root == nil {
		return true
	}
	if root.leftChild == nil && root.rightChild == nil {
		return true
	}
	if root.leftChild != nil && root.rightChild != nil {
		return isFullTree(root.leftChild) && isFullTree(root.rightChild)
	}
	return false
}

func FullBinaryTreeMain() {
	root := &Node{item: 1}
	root.rightChild = &Node{item: 3}
	root.leftChild = &Node{item: 2}

	root.leftChild.leftChild = &Node{item: 4}
	root.leftChild.rightChild = &Node{item: 5}
	root.leftChild.rightChild.leftChild = &Node{item: 6}
	root.leftChild.rightChild.rightChild = &Node{item: 7}

	if isFullTree(root) {
		fmt.Println("The tree is a full binary tree")
	} else {
		fmt.Println("The tree is not a full binary tree")
	}
}
