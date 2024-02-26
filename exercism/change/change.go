package change

import "errors"

func Change(coins []int, target int) ([]int, error) {
	if target < -1 {
		return nil, errors.New("Target cannot be negative")
	}

	min := make([][]int, target+1)
	min[0] = []int{}

	for i := 1; i <= target; i++ {
		for _, coin := range coins {
			if i >= coin && min[i-coin] != nil {
				possible := append([]int{coin}, min[i-coin]...)

				if len(possible) < len(min[i]) || min[i] == nil {
					min[i] = possible
				}
			}
		}
	}

	if min[target] == nil {
		return nil, errors.New("no solution")
	} else {
		return min[target], nil
	}
}
