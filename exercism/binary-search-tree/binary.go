package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

func (bst *BinarySearchTree) Insert(i int) {
	if bst == nil {
		return
	}

	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

func (bst *BinarySearchTree) SortedData() []int {
	if bst == nil {
		return []int{}
	}

	return append(append(bst.left.SortedData(), bst.data), bst.right.SortedData()...)
}
