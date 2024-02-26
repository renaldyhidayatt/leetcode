package dynamic

func CutRodRec(price []int, length int) int {
	if length == 0 {
		return 0
	}

	q := -1

	for i := 1; i <= length; i++ {
		q = Max(q, price[i]+CutRodRec(price, length-i))
	}

	return q
}

func CutRodDp(price []int, length int) int {
	r := make([]int, length+1)
	r[0] = 0

	for j := 1; j <= length; j++ {
		q := -1

		for i := 1; i <= j; i++ {
			q = Max(q, price[i]+r[j-i])
		}

		r[j] = q
	}

	return r[length]
}
