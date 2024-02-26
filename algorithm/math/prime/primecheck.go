package prime

func TrialDivision(n int64) bool {
	if n < 2 {
		return false
	}

	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func OptimizedTrialDivision(n int64) bool {
	if n < 2 {
		return false
	}

	if n < 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
