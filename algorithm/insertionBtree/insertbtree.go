package insertionbtree

import "fmt"

type bTreeNode struct {
	leaf  bool
	keys  []int
	child []*bTreeNode
}

type bTree struct {
	root *bTreeNode
	t    int
}

func newBTree(t int) *bTree {
	return &bTree{
		root: &bTreeNode{leaf: true},
		t:    t,
	}
}

func (b *bTree) insert(k int) {
	root := b.root
	if len(root.keys) == (2*b.t)-1 {
		temp := &bTreeNode{}
		b.root = temp
		temp.child = append(temp.child, root)
		b.splitChild(temp, 0)
		b.insertNonFull(temp, k)
	} else {
		b.insertNonFull(root, k)
	}
}

func (b *bTree) insertNonFull(x *bTreeNode, k int) {
	i := len(x.keys) - 1
	if x.leaf {
		x.keys = append(x.keys, k)
		for i >= 0 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
	} else {
		for i >= 0 && k < x.keys[i] {
			i--
		}
		i++
		if len(x.child[i].keys) == (2*b.t)-1 {
			b.splitChild(x, i)
			if k > x.keys[i] {
				i++
			}
		}
		b.insertNonFull(x.child[i], k)
	}
}

func (b *bTree) splitChild(x *bTreeNode, i int) {
	t := b.t
	y := x.child[i]
	z := &bTreeNode{leaf: y.leaf}
	x.child = append(x.child, z)
	copy(x.child[i+1:], x.child[i:])
	x.child[i+1] = z
	x.keys = append(x.keys, 0)
	copy(x.keys[i+1:], x.keys[i:])
	x.keys[i] = y.keys[t-1]
	z.keys = append(z.keys, y.keys[t:]...)
	y.keys = y.keys[:t-1]
	if !y.leaf {
		z.child = append(z.child, y.child[t:]...)
		y.child = y.child[:t-1]
	}
}

func (b *bTree) printTree(x *bTreeNode, l int) {
	fmt.Printf("Level %d %d: ", l, len(x.keys))
	for _, i := range x.keys {
		fmt.Print(i, " ")
	}
	fmt.Println()
	l++
	if len(x.child) > 0 {
		for _, i := range x.child {
			b.printTree(i, l)
		}
	}
}

func InsertBtreeMain() {
	b := newBTree(3)

	for i := 0; i < 10; i++ {
		b.insert(i)
	}

	b.printTree(b.root, 0)
}
