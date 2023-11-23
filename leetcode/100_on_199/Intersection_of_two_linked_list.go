package on199

import "fmt"

func GetIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {

		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}
