package btree

type BTreeNode struct {
	leaf bool
	keys []struct {
		key1 int
		key2 int
	}
	child []*BTreeNode
}

type BTree struct {
	root *BTreeNode
	t    int
}

func (b *BTree) insert(k int) {
	root := b.root
	if len(root.keys) == (2*b.t)-1 {
		temp := &BTreeNode{}
		b.root = temp
		temp.child = append(temp.child, root)
		b.splitChild(temp, 0)
		b.insertNonFull(temp, k)
	} else {
		b.insertNonFull(root, k)
	}
}

func (b *BTree) insertNonFull(x *BTreeNode, k int) {
	i := len(x.keys) - 1
	if x.leaf {
		x.keys = append(x.keys, struct {
			key1 int
			key2 int
		}{key1: 0, key2: 0})

		for i >= 0 && k[0] < x.keys[i].key1 {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = struct {
			key1 int
			key2 int
		}{key1: k[0], key2: k[1]}
	} else {
		for i >= 0 && k[0] < x.keys[i].key1 {
			i--
		}
		i++
		if len(x.child[i].keys) == (2*b.t)-1 {
			b.splitChild(x, i)
			if k[0] > x.keys[i].key1 {
				i++
			}
		}
		b.insertNonFull(x.child[i], k)
	}
}
