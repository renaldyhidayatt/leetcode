package concatenationofconsecutive

func concatenatedBinary(n int) int {
	res, mod, shift := 0, 1000000007, 0
	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			shift++
		}
		res = ((res << shift) + i) % mod
	}
	return res
}
