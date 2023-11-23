package binary

func ReverseBits(number uint) uint {
	result := uint(0)

	intSize := 31

	for number != 0 {
		result += (number & 1) << uint(intSize)
		number = number >> 1

		intSize -= 1
	}

	return result
}
