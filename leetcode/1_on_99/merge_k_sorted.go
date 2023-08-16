package on99

func MergeKList(lists []*ListNode) *ListNode {
	length := len(lists)

	if length < 1 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	num := length / 2

	left := MergeKList(lists[:num])
	right := MergeKList(lists[num:])

	return MergeTwoList(left, right)
}
