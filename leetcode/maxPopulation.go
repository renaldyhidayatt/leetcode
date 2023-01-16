package main

import "sort"

func maximumPopulation(logs [][]int) int {
	var year [][]int
	for _, log := range logs {
		year = append(year, []int{log[0], 1})
		year = append(year, []int{log[1], -1})
	}
	sort.Slice(year, func(i, j int) bool { return year[i][0] < year[j][0] })

	maxCnt := 0
	x := year[0][0]
	cnt := 0
	for _, y := range year {
		cnt += y[1]
		if cnt > maxCnt {
			maxCnt = cnt
			x = y[0]
		}
	}

	return x

}
