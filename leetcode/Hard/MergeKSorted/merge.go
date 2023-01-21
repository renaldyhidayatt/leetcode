package mergeksorted

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	} else if len(lists) == 0 {
		return nil
	}

	return MergeK(lists, 0, len(lists)-1)
}

func MergeK(lists []*ListNode, low, high int) *ListNode {
	if low == high {
		return lists[low]
	} else if low+1 == high {
		return MergeTwolist(lists[low], lists[high])
	}
	mid := (low + high) / 2
	return MergeTwolist(MergeK(lists, low, mid), MergeK(lists, mid+1, high))
}

func MergeTwolist(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	curr := &ListNode{-1, nil}
	head := curr
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return head.Next
}
