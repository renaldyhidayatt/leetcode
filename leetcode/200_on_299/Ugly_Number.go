package on299

func IsUgly(num int) bool {
	if num > 0 {
		for _, i := range []int{2, 3, 5} {
			for num%i == 0 {
				num /= i
			}
		}
	}

	return num == 1
}
