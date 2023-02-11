package theskylineproblem

import "sort"

const LEFTSIDE = 1
const RIGHTSIDE = 2

type Point struct {
	xAxis int
	side  int
	index int
}

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}
	if len(buildings) == 0 {
		return res
	}
	allPoints, bit := make([]Point, 0), BinaryIndexedTree{}
	for i, b := range buildings {
		allPoints = append(allPoints, Point{xAxis: b[0], side: LEFTSIDE, index: i})
		allPoints = append(allPoints, Point{xAxis: b[1], side: RIGHTSIDE, index: i})
	}
	sort.Slice(allPoints, func(i, j int) bool {
		if allPoints[i].xAxis == allPoints[j].xAxis {
			return allPoints[i].side < allPoints[j].side
		}
		return allPoints[i].xAxis < allPoints[j].xAxis
	})
	bit.Init(len(allPoints))
	kth := make(map[Point]int)
	for i := 0; i < len(allPoints); i++ {
		kth[allPoints[i]] = i
	}
	for i := 0; i < len(allPoints); i++ {
		pt := allPoints[i]
		if pt.side == LEFTSIDE {
			bit.Add(kth[Point{xAxis: buildings[pt.index][1], side: RIGHTSIDE, index: pt.index}], buildings[pt.index][2])
		}
		currHeight := bit.Query(kth[pt] + 1)
		if len(res) == 0 || res[len(res)-1][1] != currHeight {
			if len(res) > 0 && res[len(res)-1][0] == pt.xAxis {
				res[len(res)-1][1] = currHeight
			} else {
				res = append(res, []int{pt.xAxis, currHeight})
			}
		}
	}
	return res
}

type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

func (bit *BinaryIndexedTree) Init(capacity int) {
	bit.tree, bit.capacity = make([]int, capacity+1), capacity
}

func (bit *BinaryIndexedTree) Add(index int, val int) {
	for ; index > 0; index -= index & -index {
		bit.tree[index] = max(bit.tree[index], val)
	}
}

func (bit *BinaryIndexedTree) Query(index int) int {
	sum := 0
	for ; index <= bit.capacity; index += index & -index {
		sum = max(sum, bit.tree[index])
	}
	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
