package on199

func HammingWeight(num uint32) int {
	count := 0

	for num != 0 {
		num = num & (num - 1)
		count++
	}

	return count
}
