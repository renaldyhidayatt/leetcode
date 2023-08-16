package on99

func MergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2

	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = MergeTwoList(l1.Next, l2)

		return l1
	}

	l2.Next = MergeTwoList(l1, l2.Next)

	return l2
}
