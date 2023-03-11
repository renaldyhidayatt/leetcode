package beautifularrangement

func countArrangement(N int) int {
	if N == 0 {
		return 0
	}

	nums, used, p, res := make([]int, N), make([]bool, N), []int{}, [][]int{}
	for i := range nums {
		nums[i] = i + 1
	}

	generatePermutation(nums, 0, p, &res, &used)
	return len(res)
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if !(checkDivisible(nums[i], len(p)+1) || checkDivisible(len(p)+1, nums[i])) {
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}

	return
}

func checkDivisible(num, d int) bool {
	tmp := num / d

	if int(tmp)*int(d) == num {
		return true
	}

	return false
}
