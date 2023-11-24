package on299

func IsPowerOfTwo3(num int) bool {
	for num >= 2 {
		if num%2 == 0 {
			num = num / 2
		} else {
			return false
		}
	}
	return num == 1
}
