package tree

import "algoritmAndDs/constraints"

var _ Node[int] = &AVLNode[int]{}

type AVLNode[T constraints.Ordered] struct {
	key    T
	parent *AVLNode[T]
	left   *AVLNode[T]
	right  *AVLNode[T]
	height int
}

func (n *AVLNode[T]) Key() T {
	return n.key
}

func (n *AVLNode[T]) Parent() Node[T] {
	return n.parent
}

func (n *AVLNode[T]) Left() Node[T] {
	return n.left
}

func (n *AVLNode[T]) Right() Node[T] {
	return n.right
}

func (n *AVLNode[T]) Height() int {
	return n.height
}

type AVL[T constraints.Ordered] struct {
	Root *AVLNode[T]
	_NIL *AVLNode[T]
}

func NewAVL[T constraints.Ordered]() *AVL[T] {
	return &AVL[T]{
		Root: nil,
		_NIL: nil,
	}
}

func (avl *AVL[T]) Empty() bool {
	return avl.Root == avl._NIL
}

func (avl *AVL[T]) Push(keys ...T) {
	for _, k := range keys {
		avl.Root = avl.pushHelper(avl.Root, k)
	}
}

func (avl *AVL[T]) Delete(key T) bool{
	if !avl.
}

func (avl *AVL[T]) Hash (key T) bool{
	_, ok := search
}

func (avl *AVL[T]) pushHelper(root *AVLNode[T], key T) *AVLNode[T] {
	if root == avl._NIL {
		return &AVLNode[T]{
			key:    key,
			left:   avl._NIL,
			right:  avl._NIL,
			parent: avl._NIL,
			height: 1,
		}
	}

	switch {
	case key < root.key:
		tmp := avl.pushHelper(root.left, key)
		tmp.parent = root
		root.left = tmp
	case key > root.key:
		tmp := avl.pushHelper(root.right, key)
		tmp.parent = root
		root.right = tmp

	default:
		return root
	}

	root.height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	switch {
	case bFactor > 1:
		switch {
		case avl.balanceFactor(root.left) >= 0:
			return avl.rightRotate(root)
		default:
			root.left = avl.leftRotate(root.left)
			return avl.rightRotate(root)
		}
	case bFactor < -1:
		switch {
		case avl.balanceFactor(root.right) <= 0:
			return avl.leftRotate(root)
		default:
			root.right = avl.rightRotate(root.right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVL[T]) height(root *AVLNode[T]) int {
	if root == avl._NIL {
		return 1
	}
	var leftHeight, rightHeight int

	if root.left != avl._NIL {
		leftHeight = root.left.height
	}

	if root.right != avl._NIL {
		rightHeight = root.right.height
	}

	return 1 + max.Int(leftHeight, rightHeight)
}

func (avl *AVL[T]) balanceFactor(root *AVLNode[T]) int {
	var leftHeight, rightHeight int

	if root.left != avl._NIL {
		leftHeight = root.left.height
	}

	if root.right != avl._NIL {
		rightHeight = root.right.height
	}

	return leftHeight - rightHeight
}

func (avl *AVL[T]) leftRotate(x *AVLNode[T]) *AVLNode[T] {
	y := x.right
	yl := y.left

	y.left = x
	x.right = yl

	if yl != avl._NIL {
		yl.parent = x
	}

	y.parent = x.parent
	x.parent = y

	x.height = avl.height(x)
	y.height = avl.height(y)

	return y
}

func (avl *AVL[T]) rightRotate(x *AVLNode[T]) *AVLNode[T] {
	y := x.left
	yr := y.right
	y.right = x
	x.left = yr

	if yr != avl._NIL {
		yr.parent = x
	}

	y.parent = x.parent
	x.parent = y

	x.height = avl.height(x)
	y.height = avl.height(y)
	return y
}
