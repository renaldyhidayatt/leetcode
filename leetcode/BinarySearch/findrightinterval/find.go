package findrightinterval

import "sort"

func findRightInterval(intervals [][]int) []int {
	intervalList := make(intervalList, len(intervals))
	for i, v := range intervals {
		intervalList[i] = interval{interval: v, index: i}
	}

	sort.Sort(intervalList)

	out := make([]int, len(intervalList))

	for i := 0; i < len(intervalList); i++ {
		index := sort.Search(len(intervalList), func(p int) bool { return intervalList[p].interval[0] >= intervalList[i].interval[1] })
		if index == len(intervalList) {
			out[intervalList[i].index] = -1
		} else {
			out[intervalList[i].index] = intervalList[index].index
		}
	}

	return out
}

type interval struct {
	interval []int
	index    int
}

type intervalList []interval

func (in intervalList) Len() int { return len(in) }
func (in intervalList) Less(i, j int) bool {
	return in[i].interval[0] <= in[j].interval[0]
}
func (in intervalList) Swap(i, j int) { in[i], in[j] = in[j], in[i] }
