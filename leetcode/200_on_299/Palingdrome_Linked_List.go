package on299

func IsPalingDrome(head *ListNode) bool {
	slice := []int{}

	for head != nil {
		slice = append(slice, head.Val)

		head = head.Next
	}

	for i, j := 0, len(slice)-1; i < j; {
		if slice[i] != slice[j] {
			return false
		}

		i++
		j--
	}

	return true
}
