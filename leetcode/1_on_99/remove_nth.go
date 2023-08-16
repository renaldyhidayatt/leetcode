package on99

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	preSlow, slow, fast := dummyHead, head, head

	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}

		n--
		fast = fast.Next
	}

	preSlow.Next = slow.Next

	return dummyHead.Next
}
