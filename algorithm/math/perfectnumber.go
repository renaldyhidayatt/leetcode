package math

func SumProperDivisors(inNumber uint) uint {
	var res = uint(0)

	if inNumber > 1 {
		res = uint(1)
	}

	for curDivisor := uint(2); curDivisor*curDivisor <= inNumber; curDivisor++ {
		if inNumber%curDivisor == 0 {
			res += curDivisor

			if curDivisor*curDivisor != inNumber {
				res += inNumber / curDivisor
			}
		}

	}
	return res
}

func IsPerfectNumber(inNumber uint) bool {
	return inNumber > 0 && SumProperDivisors(inNumber) == inNumber
}
