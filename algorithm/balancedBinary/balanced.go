package balancedbinary

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Height struct {
	height int
}

func isHeightBalanced(root *Node, height *Height) bool {
	leftHeight := &Height{}
	rightHeight := &Height{}

	if root == nil {
		return true
	}

	l := isHeightBalanced(root.left, leftHeight)
	r := isHeightBalanced(root.right, rightHeight)

	height.height = int(math.Max(float64(leftHeight.height), float64(rightHeight.height))) + 1

	if math.Abs(float64(leftHeight.height-rightHeight.height)) <= 1 {
		return l && r
	}

	return false
}

func BalancedMain() {

	height := &Height{}

	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	if isHeightBalanced(root, height) {
		fmt.Println("The tree is balanced")
	} else {
		fmt.Println("The tree is not balanced")
	}

}
