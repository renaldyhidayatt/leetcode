package xor

func Encrypt(key byte, plaintext []byte) []byte {
	cipherText := []byte{}

	for _, ch := range plaintext {
		cipherText = append(cipherText, key^ch)
	}

	return cipherText
}

func Decrypt(key byte, cipherText []byte) []byte {
	plainText := []byte{}

	for _, ch := range cipherText {
		plainText = append(plainText, key^ch)
	}

	return plainText
}
