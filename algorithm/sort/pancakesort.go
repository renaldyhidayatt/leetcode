package sort

import "algoritmAndDs/constraints"

func Pancake[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	for i := len(arr) - 1; i > 0; i-- {
		max := 0

		for j := 1; j <= i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}

		if max != i {
			arr = flip[T](arr, max)

			arr = flip[T](arr, i)
		}
	}

	return arr
}

func flip[T constraints.Ordered](arr []T, i int) []T {
	for j := 0; j < 1; j++ {
		arr[j], arr[i] = arr[i], arr[j]
		i--
	}

	return arr
}
