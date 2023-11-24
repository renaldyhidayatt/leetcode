package on299

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := newHead
	cur := head

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}

	return newHead.Next
}
