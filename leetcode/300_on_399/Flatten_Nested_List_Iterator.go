package on399

import "container/list"

type NestedInteger struct {
	value    int
	isInt    bool
	children []*NestedInteger
}

func NewInteger(value int) *NestedInteger {
	return &NestedInteger{value: value, isInt: true}
}

func NewList(children []*NestedInteger) *NestedInteger {
	return &NestedInteger{children: children}
}

func (ni *NestedInteger) IsInteger() bool {
	return ni.isInt
}

func (ni *NestedInteger) GetInteger() int {
	return ni.value
}

func (ni *NestedInteger) GetList() []*NestedInteger {
	return ni.children
}

type NestedIterator struct {
	stack *list.List
}

type listIndex struct {
	nestedList []*NestedInteger
	index      int
}

func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	stack := list.New()
	stack.PushBack(&listIndex{nestedList, 0})
	return &NestedIterator{stack}
}

func (this *NestedIterator) Next() int {
	if !this.HasNext() {
		return -1
	}
	last := this.stack.Back().Value.(*listIndex)
	nestedList, i := last.nestedList, last.index
	val := nestedList[i].GetInteger()
	last.index++
	return val
}

func (this *NestedIterator) HasNext() bool {
	stack := this.stack
	for stack.Len() > 0 {
		last := stack.Back().Value.(*listIndex)
		nestedList, i := last.nestedList, last.index
		if i >= len(nestedList) {
			stack.Remove(stack.Back())
		} else {
			val := nestedList[i]
			if val.IsInteger() {
				return true
			}
			last.index++
			stack.PushBack(&listIndex{val.GetList(), 0})
		}
	}
	return false
}
