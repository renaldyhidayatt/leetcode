package on199

func HasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
