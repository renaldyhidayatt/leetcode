package caesar

func Encrypt(input string, key int) string {
	key8 := byte(key%26+26) % 26

	var outputBuffer []byte

	for _, b := range []byte(input) {
		newByte := b

		if 'A' <= b && b <= 'Z' {
			outputBuffer = append(outputBuffer, 'A'+(newByte-'A'+key8)%26)
		} else if 'a' <= b && b <= 'z' {
			outputBuffer = append(outputBuffer, 'a'+(newByte-'a'+key8)%26)
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}
	}

	return string(outputBuffer)
}

func Decrypt(input string, key int) string {
	return Encrypt(input, 26-key)
}
