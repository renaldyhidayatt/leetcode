package on99

func Partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	before := beforeHead
	afterHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}

		head = head.Next
	}

	after.Next = nil
	before.Next = afterHead.Next

	return beforeHead.Next
}
