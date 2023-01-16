package main

const COLOR_CNT = 26

type Node struct {
	incomming int
	color     byte
	processed bool
	colors    [COLOR_CNT]int
	children  []int
}

func (n *Node) AddMyColor() {
	n.colors[n.color]++
}

func (n *Node) Processed() {
	n.processed = true
}

func (n *Node) IsProcessed() bool {
	return n.processed
}

func (n *Node) AddIncomming(inc int) {
	n.incomming += inc
}

func (n *Node) Incomming() int {
	return n.incomming
}

func (n *Node) AddChildNode(child int) {
	n.children = append(n.children, child)
}

func (n *Node) Children() []int {
	return n.children
}

func (n *Node) ChildrenCount() int {
	return len(n.children)
}

func (n *Node) Colors() [COLOR_CNT]int {
	return n.colors
}

func (n *Node) MaxColorValue() int {
	m := -1
	for _, c := range n.colors {
		if c > m {
			m = c
		}
	}
	return m
}
