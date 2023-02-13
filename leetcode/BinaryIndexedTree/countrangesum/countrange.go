package countrangesum

type Node struct {
	val   int64
	cnt   int
	left  *Node
	right *Node
}

func NewNode(v int64) *Node {
	return &Node{val: v, cnt: 1, left: nil, right: nil}
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val int64) {
	t.InsertNode(&t.root, val)
}

func (t *Tree) InsertNode(root **Node, val int64) {
	if *root == nil {
		*root = NewNode(val)
		return
	}

	(*root).cnt++

	if val < (*root).val {
		t.InsertNode(&(*root).left, val)
	} else if val > (*root).val {
		t.InsertNode(&(*root).right, val)
	}
}

func (t *Tree) LessThan(sum int64, val int) int {
	return t.LessThanNode(t.root, sum, val, 0)
}

func (t *Tree) LessThanNode(root *Node, sum int64, val int, res int) int {
	if root == nil {
		return res
	}

	if sum-root.val < int64(val) {
		res += root.cnt - t.getCount(root.left)
		return t.LessThanNode(root.left, sum, val, res)
	} else if sum-root.val > int64(val) {
		return t.LessThanNode(root.right, sum, val, res)
	} else {
		return res + t.getCount(root.right)
	}
}

func (t *Tree) getCount(root *Node) int {
	if root == nil {
		return 0
	}
	return root.cnt
}

func countRangeSum(nums []int, lower int, upper int) int {
	tree := Tree{}
	tree.Insert(0)
	sum, res := int64(0), 0

	for _, n := range nums {
		sum += int64(n)
		lcnt := tree.LessThan(sum, lower)
		hcnt := tree.LessThan(sum, upper+1)
		res += hcnt - lcnt
		tree.Insert(sum)
	}

	return res
}
