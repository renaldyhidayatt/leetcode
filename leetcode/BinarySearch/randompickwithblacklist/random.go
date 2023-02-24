package randompickwithblacklist

import "math/rand"

type Solution struct {
	M        int
	BlackMap map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	blackMap := map[int]int{}

	for i := 0; i < len(blacklist); i++ {
		blackMap[blacklist[i]] = 1
	}

	M := N - len(blacklist)

	for _, value := range blacklist {
		if value < M {
			for {
				if _, ok := blackMap[N-1]; ok {
					N--
				} else {
					break
				}
			}
			blackMap[value] = N - 1

			N--
		}
	}

	return Solution{BlackMap: blackMap, M: M}
}

func (this *Solution) Pick() int {
	idx := rand.Intn(this.M)

	if _, ok := this.BlackMap[idx]; ok {
		return this.BlackMap[idx]
	}

	return idx
}
