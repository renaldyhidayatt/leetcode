package math

func AliquotSum(n int) (int, error) {
	switch {
	case n < 0:
		return 0, ErrPosArgsOnly
	case n == 0:
		return 0, ErrNonZeroArgsOnly

	default:
		var sum int

		for i := 1; i <= n/2; i++ {
			if n%i == 0 {
				sum += i
			}
		}

		return sum, nil
	}
}
