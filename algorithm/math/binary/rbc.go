package binary

func SequenceGrayCode(n uint) []uint {
	result := make([]uint, 0)
	var i uint
	for i = 0; i < 1<<n; i++ {
		result = append(result, i^(i>>1))
	}
	return result
}
