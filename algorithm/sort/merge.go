package sort

import (
	"algoritmAndDs/constraints"
	"algoritmAndDs/math/min"
	"sync"
)

func merge[T constraints.Ordered](a []T, b []T) []T {
	var r = make([]T, len(a)+len(b))

	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func Merge[T constraints.Ordered](items []T) []T {
	if len(items) <= 1 {
		return items
	}

	var middle = len(items) / 2
	var a = Merge(items[:middle])
	var b = Merge(items[middle:])
	return merge(a, b)
}

func MergeIter[T constraints.Ordered](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			temp := merge[T](items[i:i+step], items[i+step:min.Int(i+2*step, len(items))])
			copy(items[i:], temp)
		}
	}

	return items
}

func ParallelMerge[T constraints.Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	if len(items) < 2048 {
		return Merge[T](items)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	var middle = len(items) / 2

	var a []T
	go func() {
		defer wg.Done()
		a = ParallelMerge[T](items[:middle])
	}()

	var b = ParallelMerge[T](items[middle:])

	wg.Wait()

	return merge[T](a, b)
}
