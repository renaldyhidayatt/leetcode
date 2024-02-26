package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}

	sieve := make([]bool, limit+1)
	var primes []int

	for i := 2; i <= limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)

			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}

	return primes
}
