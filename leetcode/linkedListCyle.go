package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCyle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next.Next
	slow := head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}
