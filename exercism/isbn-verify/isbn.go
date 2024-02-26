package isbnverify

import "strings"

func IsValidISBN(isbnWithHypen string) bool {
	isbn := strings.ReplaceAll(isbnWithHypen, "-", "")

	if len(isbn) != 10 {
		return false
	}

	sum := 0

	multiplicator := 10

	for i := 0; i < len(isbn); i++ {
		digit := int(isbn[i] - '0')

		if !(digit >= 0 && digit <= 9) {
			if i == len(isbn)-1 && isbn[i] == 'X' {
				digit = 10
			} else {
				return false
			}
		}

		sum += digit * multiplicator

		multiplicator--
	}

	return sum%11 == 0

}
