package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	var cipher []rune

	for _, a := range plain {
		if unicode.IsUpper(a) {
			a = rune('A' + (int(a-'A')+shiftKey)%26)
		} else if unicode.IsLower(a) {
			a = rune('a' + (int(a-'a')+shiftKey)%26)
		}

		cipher = append(cipher, a)
	}

	return string(cipher)
}
