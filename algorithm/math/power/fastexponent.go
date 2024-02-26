package power

func IterativePower(n uint, power uint) uint {
	var res uint = 1

	for power > 0 {
		if (power & 1) != 0 {
			res = res * n
		}

		power = power >> 1
		n *= n
	}

	return res
}

func ResursivePower(n uint, power uint) uint {
	if power == 0 {
		return 1
	}

	var temp = ResursivePower(n, power/2)

	if power%2 == 0 {
		return temp * temp
	}

	return n * temp * temp
}

func ResursivePower1(n uint, power uint) uint {
	if power == 0 {
		return 1
	} else if power%2 == 0 {
		return ResursivePower1(n, power/2) * ResursivePower1(n, power/2)
	} else {
		return n * ResursivePower1(n, power/2) * ResursivePower1(n, power/2)
	}
}
