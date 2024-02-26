package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	list := map[int]bool{}
	sum := 0

	for _, i := range divisors {
		if i == 0 {
			continue
		}

		for j := i; j < limit; j += i {
			if _, ok := list[j]; !ok {
				list[j] = true

				sum += j
			}
		}
	}
	return sum
}
