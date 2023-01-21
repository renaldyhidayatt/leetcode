package removenthnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast, curr := head, head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// n == len
	if fast == nil {
		head = head.Next
		return head
	}
	// Move both pointers, until reach tail
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	curr = slow.Next
	slow.Next = curr.Next
	return head
}

func RunningRemoveFromEnd() {
	l := &ListNode{
		Val:  10,
		Next: &ListNode{},
	}

	fmt.Println(removeNthFromEnd(l, 3))
}
