package romantointeger

import "fmt"

func romantToInt(s string) int {
	roman := map[rune]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	prev, total := 0, 0

	for _, c := range s {
		curr := roman[c]
		total += curr
		if curr > prev {
			total -= 2 * prev
		}

		prev = curr
	}

	return total
}

func RunningRomanToInt() {
	fmt.Println(romantToInt("III"))
}
