package binarytree

import "fmt"

type Node struct {
	left  *Node
	right *Node
	val   int
}

func (n *Node) TraversePreOrder() {
	fmt.Print(n.val, " ")
	if n.left != nil {
		n.left.TraversePreOrder()
	}
	if n.right != nil {
		n.right.TraversePreOrder()
	}
}

func (n *Node) TraverseInOrder() {
	if n.left != nil {
		n.left.TraverseInOrder()
	}
	fmt.Print(n.val, " ")
	if n.right != nil {
		n.right.TraverseInOrder()
	}
}

func (n *Node) TraversePostOrder() {
	if n.left != nil {
		n.left.TraversePostOrder()
	}
	if n.right != nil {
		n.right.TraversePostOrder()
	}
	fmt.Print(n.val, " ")
}

func BinaryTreeMain() {
	root := &Node{val: 1}

	root.left = &Node{val: 2}
	root.right = &Node{val: 3}

	root.left.left = &Node{val: 4}

	fmt.Print("Pre order Traversal: ")
	root.TraversePreOrder()
	fmt.Print("\nIn order Traversal: ")
	root.TraverseInOrder()
	fmt.Print("\nPost order Traversal: ")
	root.TraversePostOrder()

}
