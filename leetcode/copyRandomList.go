package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	p := head

	for p != nil {

		next := p.Next
		copy := &Node{Val: p.Val, Next: nil, Random: nil}
		p.Next = copy
		copy.Next = next
		p = next
	}

	p = head

	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	p = head

	var headCopy *Node
	if p != nil {
		headCopy = p.Next
	} else {
		headCopy = nil
	}

	for p != nil {
		copy := p.Next
		p.Next = copy.Next
		p = p.Next
		if p != nil {
			copy.Next = p.Next
		}
	}

	return headCopy
}
