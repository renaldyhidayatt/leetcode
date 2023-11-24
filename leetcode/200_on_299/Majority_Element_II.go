package on299

func MajorityElement229(nums []int) []int {
	result, m := make([]int, 0), make(map[int]int)

	for _, val := range nums {
		if v, ok := m[val]; ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}
	}

	for k, v := range m {
		if v > len(nums)/3 {
			result = append(result, k)
		}
	}

	return result
}
