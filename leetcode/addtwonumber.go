package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{0, nil}
	curr := head
	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{val % 10, nil}
		curr = curr.Next
		carry = val / 10
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return head.Next
}
