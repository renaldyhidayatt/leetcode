package on399

import (
	"math/rand"
)

type Solution struct {
	head *ListNode
}

func Constructor382(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	scope, selectPoint, curr := 1, 0, this.head

	for curr != nil {
		if rand.Float64() < 1.0/float64(scope) {
			selectPoint = curr.Val
		}

		scope += 1

		curr = curr.Next
	}

	return selectPoint
}
