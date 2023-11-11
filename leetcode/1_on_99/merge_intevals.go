package on99

type Interval struct {
	Start int
	End   int
}

func Merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	quickSort(intervals, 0, len(intervals)-1)
	res := make([]Interval, 0)
	res = append(res, intervals[0])
	curIndex := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > res[curIndex].End {
			curIndex++
			res = append(res, intervals[i])
		} else {
			res[curIndex].End = max(intervals[i].End, res[curIndex].End)
		}
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func partitionSort(a []Interval, lo, h1 int) int {
	pivot := a[h1]
	i := lo - 1

	for j := lo; j < h1; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[j], a[i] = a[i], a[j]
		}

	}
	a[i+1], a[h1] = a[h1], a[i+1]

	return i + 1
}

func quickSort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
