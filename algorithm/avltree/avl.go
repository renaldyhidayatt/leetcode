package avltree

import "fmt"

type TreeNode struct {
	key    int
	left   *TreeNode
	right  *TreeNode
	height int
}

type AVLTree struct{}

// Insert a node into the tree
func (t *AVLTree) insertNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{key: key, height: 1}
	}
	if key < root.key {
		root.left = t.insertNode(root.left, key)
	} else {
		root.right = t.insertNode(root.right, key)
	}

	root.height = 1 + max(t.getHeight(root.left), t.getHeight(root.right))

	balanceFactor := t.getBalance(root)
	if balanceFactor > 1 {
		if key < root.left.key {
			return t.rightRotate(root)
		} else {
			root.left = t.leftRotate(root.left)
			return t.rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if key > root.right.key {
			return t.leftRotate(root)
		} else {
			root.right = t.rightRotate(root.right)
			return t.leftRotate(root)
		}
	}

	return root
}

// Delete a node from the tree
func (t *AVLTree) deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key < root.key {
		root.left = t.deleteNode(root.left, key)
	} else if key > root.key {
		root.right = t.deleteNode(root.right, key)
	} else {
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		} else if root.right == nil {
			temp := root.left
			root = nil
			return temp
		}
		temp := t.getMinValueNode(root.right)
		root.key = temp.key
		root.right = t.deleteNode(root.right, temp.key)
	}
	if root == nil {
		return root
	}

	root.height = 1 + max(t.getHeight(root.left), t.getHeight(root.right))

	balanceFactor := t.getBalance(root)
	if balanceFactor > 1 {
		if t.getBalance(root.left) >= 0 {
			return t.rightRotate(root)
		} else {
			root.left = t.leftRotate(root.left)
			return t.rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if t.getBalance(root.right) <= 0 {
			return t.leftRotate(root)
		} else {
			root.right = t.rightRotate(root.right)
			return t.leftRotate(root)
		}
	}

	return root
}

func (t *AVLTree) leftRotate(z *TreeNode) *TreeNode {
	y := z.right
	T2 := y.left

	y.left = z
	z.right = T2

	z.height = 1 + max(t.getHeight(z.left), t.getHeight(z.right))
	y.height = 1 + max(t.getHeight(y.left), t.getHeight(y.right))

	return y
}

func (t *AVLTree) rightRotate(z *TreeNode) *TreeNode {
	y := z.left
	T3 := y.right

	y.right = z
	z.left = T3

	z.height = 1 + max(t.getHeight(z.left), t.getHeight(z.right))
	y.height = 1 + max(t.getHeight(y.left), t.getHeight(y.right))

	return y
}

func (t *AVLTree) getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.height
}

// Get the balance factor of a node
func (t *AVLTree) getBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return t.getHeight(root.left) - t.getHeight(root.right)
}

// Get the minimum value node in the tree
func (t *AVLTree) getMinValueNode(root *TreeNode) *TreeNode {
	if root == nil || root.left == nil {
		return root
	}
	return t.getMinValueNode(root.left)
}

// Perform an inorder traversal of the tree (visit left subtree, root, then right subtree)
func (t *AVLTree) inorder(root *TreeNode) {
	if root != nil {
		t.inorder(root.left)
		fmt.Printf("%d ", root.key)
		t.inorder(root.right)
	}
}

// Perform a preorder traversal of the tree (visit root, left subtree, then right subtree)
func (t *AVLTree) preorder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.key)
		t.preorder(root.left)
		t.preorder(root.right)
	}
}

// Perform a postorder traversal of the tree (visit left subtree, right subtree, then root)
func (t *AVLTree) postorder(root *TreeNode) {
	if root != nil {
		t.postorder(root.left)
		t.postorder(root.right)
		fmt.Printf("%d ", root.key)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func AvlTreeMain() {
	tree := &AVLTree{}
	root := tree.insertNode(nil, 10)
	root = tree.insertNode(root, 20)
	root = tree.insertNode(root, 30)
	root = tree.insertNode(root, 40)
	root = tree.insertNode(root, 50)
	root = tree.insertNode(root, 25)
	fmt.Println("Preorder traversal after insertion:")
	tree.preorder(root)
	fmt.Println()

	root = tree.deleteNode(root, 30)

	fmt.Println("Preorder traversal after deletion:")
	tree.preorder(root)
	fmt.Println()
}
