package randompoint

import "math/rand"

type Solution struct {
	rects [][]int
	arr   []int
}

func Constructor497(rects [][]int) Solution {
	s := Solution{
		rects: rects,
		arr:   make([]int, len(rects)),
	}

	for i := 0; i < len(rects); i++ {
		area := (rects[i][2] - rects[i][0] + 1) * (rects[i][3] - rects[i][1] + 1)
		if area < 0 {
			area = -area
		}

		if i == 0 {
			s.arr[0] = area
		} else {
			s.arr[i] = s.arr[i-1] + area
		}
	}
	return s
}

func (so *Solution) Pick() []int {
	r := rand.Int() % so.arr[len(so.arr)-1]

	low, high, index := 0, len(so.arr)-1, -1

	for low <= high {
		mid := low + (high-low)>>1

		if so.arr[mid] > r {
			if mid == 0 || so.arr[mid-1] <= r {
				index = mid
				break
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if index == -1 {
		index = low
	}

	if index > 0 {
		r = r - so.arr[index-1]
	}
	length := so.rects[index][2] - so.rects[index][0]
	return []int{so.rects[index][0] + r%(length+1), so.rects[index][1] + r/(length+1)}
}
