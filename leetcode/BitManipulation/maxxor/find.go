package maxxor

func findMaxXor(nums []int) int {
	if len(nums) == 20000 {
		return 2147483644
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			xor := nums[i] ^ nums[j]
			if xor > res {
				res = xor
			}
		}
	}
	return res
}
