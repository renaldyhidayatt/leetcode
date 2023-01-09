package binarysearchtree

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

// Inorder traversal
func inorder(root *Node) {
	if root != nil {
		// Traverse left
		inorder(root.left)

		// Traverse root
		fmt.Printf("%d -> ", root.key)

		// Traverse right
		inorder(root.right)
	}
}

// Insert a node
func insert(node *Node, key int) *Node {
	// Return a new node if the tree is empty
	if node == nil {
		return &Node{key: key}
	}

	// Traverse to the right place and insert the node
	if key < node.key {
		node.left = insert(node.left, key)
	} else {
		node.right = insert(node.right, key)
	}

	return node
}

// Find the inorder successor
func minValueNode(node *Node) *Node {
	current := node

	// Find the leftmost leaf
	for current.left != nil {
		current = current.left
	}

	return current
}

// Deleting a node
func deleteNode(root *Node, key int) *Node {
	// Return if the tree is empty
	if root == nil {
		return root
	}

	// Find the node to be deleted
	if key < root.key {
		root.left = deleteNode(root.left, key)
	} else if key > root.key {
		root.right = deleteNode(root.right, key)
	} else {
		// If the node is with only one child or no child
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		} else if root.right == nil {
			temp := root.left
			root = nil
			return temp
		}

		// If the node has two children,
		// place the inorder successor in position of the node to be deleted
		temp := minValueNode(root.right)

		root.key = temp.key

		// Delete the inorder successor
		root.right = deleteNode(root.right, temp.key)
	}

	return root
}

func BinarySearchTreeMain() {
	var root *Node

	root = insert(root, 8)
	root = insert(root, 3)
	root = insert(root, 1)
	root = insert(root, 6)
	root = insert(root, 7)
	root = insert(root, 10)
	root = insert(root, 14)
	root = insert(root, 4)

	fmt.Print("Inorder traversal: ")
	inorder(root)

	fmt.Println("\nDelete 10")
	root = deleteNode(root, 10)
	fmt.Print("Inorder traversal: ")
	inorder(root)
}
