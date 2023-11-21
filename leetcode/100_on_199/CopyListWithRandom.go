package on199

type NodeCopy struct {
	Val    int
	Next   *NodeCopy
	Random *NodeCopy
}

func CopyRandomList(head *NodeCopy) *NodeCopy {
	if head == nil {
		return nil
	}

	tempHead := copyNodeToLinkedList(head)
	return splitLinkedList(tempHead)
}

func splitLinkedList(head *NodeCopy) *NodeCopy {
	cur := head
	head = head.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}
	return head
}

func copyNodeToLinkedList(head *NodeCopy) *NodeCopy {
	cur := head
	for cur != nil {
		node := &NodeCopy{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next, cur = node, cur.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	return head
}
