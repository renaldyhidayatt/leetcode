package on199

func ReorderList(head *ListNode) *ListNode {
	array := listToArray(head)
	length := len(array)
	if length == 0 {
		return head
	}
	cur := head
	last := head
	for i := 0; i < len(array)/2; i++ {
		tmp := &ListNode{Val: array[length-1-i], Next: cur.Next}
		cur.Next = tmp
		cur = tmp.Next
		last = tmp
	}
	if length%2 == 0 {
		last.Next = nil
	} else {
		cur.Next = nil
	}
	return head
}

func listToArray(head *ListNode) []int {
	array := []int{}
	if head == nil {
		return array
	}
	cur := head
	for cur != nil {
		array = append(array, cur.Val)
		cur = cur.Next
	}
	return array
}
