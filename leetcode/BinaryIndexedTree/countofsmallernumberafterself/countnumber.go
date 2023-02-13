package countofsmallernumberafterself

import "fmt"

type BinarySearchTreeNode struct {
	val   int
	less  int
	count int
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{val: value, less: 0, count: 1, left: nil, right: nil}
}

func (bstn *BinarySearchTreeNode) Insert(value int, numLessThan *int) {
	if value < bstn.val {
		bstn.less++
		if bstn.left == nil {
			bstn.left = NewBinarySearchTreeNode(value)
		} else {
			bstn.left.Insert(value, numLessThan)
		}
	} else if value > bstn.val {
		*numLessThan += bstn.less + bstn.count

		if bstn.right == nil {
			bstn.right = NewBinarySearchTreeNode(value)
		} else {
			bstn.right.Insert(value, numLessThan)
		}
	} else {
		*numLessThan += bstn.less
		bstn.count++
	}
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func NewBinarySearchTree(value int) *BinarySearchTree {
	return &BinarySearchTree{NewBinarySearchTreeNode(value)}
}

func (bst *BinarySearchTree) FreeTree(root *BinarySearchTreeNode) {
	if root == nil {
		return
	}
	if root.left != nil {
		bst.FreeTree(root.left)
	}
	if root.right != nil {
		bst.FreeTree(root.right)
	}
	root = nil
}

func (bst *BinarySearchTree) Insert(value int, numLessThan *int) {
	bst.root.Insert(value, numLessThan)
}

type Solution struct{}

func (s *Solution) countSmaller(nums []int) []int {
	counts := make([]int, len(nums))
	if len(nums) == 0 {
		return counts
	}
	tree := NewBinarySearchTree(nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		numLessThan := 0
		tree.Insert(nums[i], &numLessThan)
		counts[i] = numLessThan
	}
	return counts
}

func main() {
	s := Solution{}
	nums := []int{5, 2, 6, 1}
	counts := s.countSmaller(nums)
	fmt.Println(counts) // Output: [2 1 1 0]
}
