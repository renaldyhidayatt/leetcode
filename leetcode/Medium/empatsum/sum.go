package empatsum

import (
	"fmt"
	"sort"
)

func threeSum(num []int, target int) [][]int {
	var result [][]int
	sort.Ints(num)

	n := len(num)

	for i := 0; i < n-2; i++ {
		if i > 0 && num[i-1] == num[i] {
			continue
		}
		a := num[i]
		low := i + 1
		high := n - 1
		for low < high {
			b := num[low]
			c := num[high]
			if a+b+c == target {
				v := []int{a, b, c}
				result = append(result, v)
				for low < n && num[low] == num[low+1] {
					low++
				}
				for high > 0 && num[high] == num[high-1] {
					high--
				}
				low++
				high--
			} else if a+b+c > target {
				for high > 0 && num[high] == num[high-1] {
					high--
				}
				high--
			} else {
				for low < n && num[low] == num[low+1] {
					low++
				}
				low++
			}
		}
	}
	return result
}

func fourSum(num []int, target int) [][]int {
	var result [][]int
	if len(num) < 4 {
		return result
	}
	sort.Ints(num)

	for i := 0; i < len(num)-3; i++ {
		if i > 0 && num[i-1] == num[i] {
			continue
		}
		n := num[i+1:]
		ret := threeSum(n, target-num[i])
		for j := 0; j < len(ret); j++ {
			ret[j] = append([]int{num[i]}, ret[j]...)
			result = append(result, ret[j])
		}
	}

	return result
}

func RunningEmpatSum() {
	num := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(num, target)
	fmt.Println(result)
}
