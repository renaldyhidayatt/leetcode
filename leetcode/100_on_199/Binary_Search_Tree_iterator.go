package on199

import "container/heap"

type BSTIterator struct {
	pq    PriorityQueueOfInt
	count int
}

func (this *BSTIterator) Next() int {
	this.count--
	return heap.Pop(&this.pq).(int)
}

func (this *BSTIterator) HasNext() bool {
	return this.count != 0
}

func Constructor173(root *TreeNode) BSTIterator {
	result, pq := []int{}, PriorityQueueOfInt{}
	postOrder173(root, &result)

	for _, v := range result {
		heap.Push(&pq, v)

	}
	bs := BSTIterator{pq: pq, count: len(result)}

	return bs
}

func postOrder173(root *TreeNode, output *[]int) {

}

type PriorityQueueOfInt []int

func (pq PriorityQueueOfInt) Len() int {
	return len(pq)
}

func (pq PriorityQueueOfInt) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueueOfInt) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueOfInt) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueueOfInt) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
