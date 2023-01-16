package main

import (
	"sort"
)

type Vector []int

func (v Vector) Len() int {
	return len(v)
}

func (v Vector) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v Vector) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

type Vector2D [][]int

func (v Vector2D) Len() int {
	return len(v)
}

func (v Vector2D) Less(i, j int) bool {
	return v[i][0] < v[j][0]
}

func (v Vector2D) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

type Interval struct {
	Length int
	End    int
}

type IntervalSet []Interval

func (s IntervalSet) Len() int {
	return len(s)
}

func (s IntervalSet) Less(i, j int) bool {
	if s[i].Length != s[j].Length {
		return s[i].Length < s[j].Length
	}
	return s[i].End < s[j].End
}

func (s IntervalSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func minInterval(intervals [][]int, queries []int) []int {
	s := make(IntervalSet, 0)

	iQueries := make(Vector2D, len(queries))
	for i := 0; i < len(queries); i++ {
		iQueries[i] = []int{queries[i], i}
	}

	sort.Sort(Vector2D(intervals))
	sort.Sort(Vector2D(iQueries))

	result := make([]int, len(queries))
	for i := range result {
		result[i] = -1
	}

	i := 0
	len := len(intervals)

	for _, iq := range iQueries {
		q := iq[0]
		idx := iq[1]

		for i < len && intervals[i][0] <= q {
			s = append(s, Interval{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]})
			i++
		}

		for len(s) > 0 {
			if s[0].End >= q {
				break
			}
			s = s[1:]
		}

		if len(s) > 0 {
			result[idx] = s[0].Length
		}
	}

	return result
}
