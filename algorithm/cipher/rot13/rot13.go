package rot13

import "algoritmAndDs/cipher/caesar"

func Rot13(input string) string {
	return caesar.Encrypt(input, 13)
}
