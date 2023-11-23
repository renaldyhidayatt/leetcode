package on199

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{Val: 0, Next: nil}

	cur, pre := head, newHead

	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}

		cur.Next = pre.Next
		pre.Next = cur
		pre = newHead
		cur = next
	}

	return newHead.Next
}
