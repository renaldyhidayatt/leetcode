package main

const MAX_NUM = 100001

func sumOfFlooredPairs(nums []int) int {
	cnt := make([]int, MAX_NUM)
	maxn := 0
	for _, n := range nums {
		cnt[n]++
		if n > maxn {
			maxn = n
		}
	}

	stats := [][]int{}
	for i := 1; i < MAX_NUM; i++ {
		if cnt[i] > 0 {
			stats = append(stats, []int{i, cnt[i]})
		}
		cnt[i] += cnt[i-1]
	}

	const MOD = 1e9 + 7
	result := 0
	for i := 0; i < len(stats); i++ {
		n := stats[i][0]
		c := stats[i][1]

		for x := 2; x <= (maxn/n)+1; x++ {
			pre := (x-1)*n - 1
			cur := x*n - 1
			if cur > MAX_NUM-1 {
				cur = MAX_NUM - 1
			}
			result = (result + (cnt[cur]-cnt[pre])*(x-1)*c) % MOD
		}
	}
	return result
}
