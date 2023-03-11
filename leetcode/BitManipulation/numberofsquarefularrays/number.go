package numberofsquarefularrays

import (
	"math"
	"sort"
)

func numSquarefulPerms(A []int) int {
	if len(A) == 0 {
		return 0
	}

	used, p, res := make([]bool, len(A)), []int{}, [][]int{}

	sort.Ints(A)
	generatePermutation(A, 0, p, &res, &used)
	return len(res)
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		checkSquareful := true
		for i := 0; i < len(p)-1; i++ {
			if !checkSquare(p[i] + p[i+1]) {
				checkSquareful = false
				break
			}
		}
		if checkSquareful {
			temp := make([]int, len(p))
			copy(temp, p)
			*res = append(*res, temp)
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			if len(p) > 0 && !checkSquare(nums[i]+p[len(p)-1]) {
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func checkSquare(num int) bool {
	tmp := math.Sqrt(float64(num))

	if int(tmp)*int(tmp) == num {
		return true
	}

	return false
}
