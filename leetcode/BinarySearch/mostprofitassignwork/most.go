package mostprofitassignwork

import (
	"sort"
)

type Task struct {
	Difficulty int
	Profit     int
}

type SortByDiff []Task

func (s SortByDiff) Len() int {
	return len(s)
}

func (s SortByDiff) Less(i, j int) bool {
	return s[i].Difficulty < s[j].Difficulty
}

func (s SortByDiff) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	if len(difficulty) == 0 || len(profit) == 0 || len(worker) == 0 {
		return 0
	}
	tasks, res, index := []Task{}, 0, 0
	for i := 0; i < len(difficulty); i++ {
		tasks = append(tasks, Task{Difficulty: difficulty[i], Profit: profit[i]})
	}
	sort.Sort(SortByDiff(tasks))
	sort.Ints(worker)
	for i := 1; i < len(tasks); i++ {
		tasks[i].Profit = max(tasks[i].Profit, tasks[i-1].Profit)
	}
	for _, w := range worker {
		for index < len(difficulty) && w >= tasks[index].Difficulty {
			index++
		}
		if index > 0 {
			res += tasks[index-1].Profit
		}
	}
	return res
}
