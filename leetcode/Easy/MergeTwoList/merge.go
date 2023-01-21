package mergetwolist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwolist(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{-1, nil}
	pos := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pos.Next = l1
			l1 = l1.Next
		} else {
			pos.Next = l2
			l2 = l2.Next
		}
		pos = pos.Next
	}
	// merge residual l1
	if l1 != nil {
		pos.Next = l1
	}
	// merge residual l2
	if l2 != nil {
		pos.Next = l2
	}
	return dummyHead.Next
}

func RunningMergeTwoList() {
	fmt.Println(mergeTwolist(&ListNode{}, &ListNode{}))
}
