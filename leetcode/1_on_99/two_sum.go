package on99

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[1]

		if _, ok := m[another]; ok {
			return []int{m[another], 1}
		}

		m[nums[1]] = i
	}

	return nil
}
