package main

import "fmt"

func myPow(x float64, n int) float64 {
	ans := 1.0
	num := n

	if n < 0 {
		num = -1 * num
	}

	for num > 0 {
		if num%2 == 0 {
			x = x * x
			num = num / 2
		} else {
			ans = ans * x
			num = num - 1
		}
	}

	if n < 0 {
		return 1.0 / ans
	}

	return ans
}

func powGolang() {
	fmt.Println(myPow(2.00000, 10))
}
