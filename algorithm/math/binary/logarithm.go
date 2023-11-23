package binary

func LogBase2(n uint32) uint32 {
	base := [5]uint32{0x2, 0xC, 0xF0, 0xFF00, 0xFFFF0000}
	exponents := [5]uint32{1, 2, 4, 8, 16}
	var result uint32
	for i := 4; i >= 0; i-- {
		if n&base[i] != 0 {
			n >>= exponents[i]
			result |= exponents[i]
		}
	}
	return result
}
