package fibonancihead

import (
	"fmt"
	"math"
)

type FibonacciTree struct {
	value int
	child []*FibonacciTree
	order int
}

func NewFibonacciTree(value int) *FibonacciTree {
	return &FibonacciTree{
		value: value,
		child: []*FibonacciTree{},
		order: 0,
	}
}

func (t *FibonacciTree) AddAtEnd(newTree *FibonacciTree) {
	t.child = append(t.child, newTree)
	t.order++
}

type FibonacciHeap struct {
	trees []*FibonacciTree
	least *FibonacciTree
	count int
}

func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{
		trees: []*FibonacciTree{},
		least: nil,
		count: 0,
	}
}

func (h *FibonacciHeap) InsertNode(value int) {
	newTree := NewFibonacciTree(value)
	h.trees = append(h.trees, newTree)
	if h.least == nil || value < h.least.value {
		h.least = newTree
	}
	h.count++
}

func (h *FibonacciHeap) GetMin() int {
	if h.least == nil {
		return 0
	}
	return h.least.value
}

func (h *FibonacciHeap) ExtractMin() int {
	smallest := h.least
	if smallest != nil {
		for _, child := range smallest.child {
			h.trees = append(h.trees, child)
		}
		for i, t := range h.trees {
			if t == smallest {
				h.trees = append(h.trees[:i], h.trees[i+1:]...)
				break
			}
		}
		if len(h.trees) == 0 {
			h.least = nil
		} else {
			h.least = h.trees[0]
			h.Consolidate()
		}
		h.count--
		return smallest.value
	}
	return 0
}

func (h *FibonacciHeap) Consolidate() {
	aux := make([]*FibonacciTree, int(math.Floor(math.Log2(float64(h.count)))+1))

	for len(h.trees) > 0 {
		x := h.trees[0]
		order := x.order
		h.trees = h.trees[1:]
		for aux[order] != nil {
			y := aux[order]
			if x.value > y.value {
				x, y = y, x
			}
			x.AddAtEnd(y)
			aux[order] = nil
			order++
		}
		aux[order] = x
	}

	h.least = nil
	for _, k := range aux {
		if k != nil {
			h.trees = append(h.trees, k)
			if h.least == nil || k.value < h.least.value {
				h.least = k
			}
		}
	}
}

func floorLog(x int) int {
	_, exponent := math.Frexp(float64(x))

	return exponent - 1
}

func FibonanciMain() {
	fibonacciHeap := NewFibonacciHeap()

	fibonacciHeap.InsertNode(7)
	fibonacciHeap.InsertNode(3)
	fibonacciHeap.InsertNode(17)
	fibonacciHeap.InsertNode(24)

	fmt.Printf("the minimum value of the fibonacci heap: %d\n", fibonacciHeap.GetMin())

	fmt.Printf("the minimum value removed: %d\n", fibonacciHeap.ExtractMin())
}
