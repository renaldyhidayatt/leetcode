package findradius

import "math"

func findRadius1(houses []int, heaters []int) int {
	res := 0
	for i := 0; i < len(houses); i++ {
		dis := math.MaxInt64
		for j := 0; j < len(heaters); j++ {
			dis = min(dis, abs(houses[i]-heaters[j]))
		}

		res = max(res, dis)
	}

	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
