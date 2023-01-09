package treetraversal

import "fmt"

type Node struct {
	left  *Node
	right *Node
	val   int
}

func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Printf("%d->", root.val)
		inorder(root.right)
	}
}

func postorder(root *Node) {
	if root != nil {
		postorder(root.left)
		postorder(root.right)
		fmt.Printf("%d->", root.val)
	}
}

func preorder(root *Node) {
	if root != nil {
		fmt.Printf("%d->", root.val)
		preorder(root.left)
		preorder(root.right)
	}
}

func TreetraversalMain() {
	root := &Node{val: 1}
	root.left = &Node{val: 2}
	root.right = &Node{val: 3}
	root.left.left = &Node{val: 4}
	root.left.right = &Node{val: 5}

	fmt.Println("Inorder traversal")
	inorder(root)

	fmt.Println("\nPreorder traversal")
	preorder(root)

	fmt.Println("\nPostorder traversal")
	postorder(root)
}
