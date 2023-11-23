package on199

func TwoSum167(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}

	return nil
}
