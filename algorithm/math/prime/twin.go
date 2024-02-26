package prime

func Twin(n int) (int, bool) {
	if OptimizedTrialDivision(int64(n)) && OptimizedTrialDivision(int64(n+2)) {
		return n + 2, true
	}
	return -1, false
}
