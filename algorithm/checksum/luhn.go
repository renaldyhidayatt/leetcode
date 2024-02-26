package checksum

func Luhn(s []byte) bool {
	n := len(s)
	number := 0
	result := 0

	for i := 0; i < n; i++ {
		number = int(s[i]) - '0'

		if i%2 != 0 {
			result += number
			continue
		}

		number *= 2

		if number > 9 {
			number -= 9
		}

		result += number
	}

	return result%10 == 0
}
