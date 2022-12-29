package main

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{-1, nil}
	dummyHead.Next = head
	prev, p := dummyHead, head
	for p != nil && p.Next != nil {
		q, r := p.Next, p.Next.Next
		prev.Next = q
		q.Next = p
		p.Next = r
		prev = p
		p = r
	}
	return dummyHead.Next
}
