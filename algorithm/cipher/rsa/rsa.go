package rsa

import (
	"algoritmAndDs/math/modular"
	"errors"
)

var ErrorFailedToEncrypt = errors.New("failed to Encrypt")

var ErrorFailedToDecrypt = errors.New("failed to Decrypt")

func Encrypt(message []rune, publicExponent, modulus int64) ([]rune, error) {
	var encrypted []rune

	for _, letter := range message {
		encryptedLetter, err := modular.Exponentiation(int64(letter), publicExponent, modulus)
		if err != nil {
			return nil, ErrorFailedToEncrypt
		}
		encrypted = append(encrypted, rune(encryptedLetter))
	}

	return encrypted, nil
}

func Decrypt(encrypted []rune, privateExponent, modulus int64) (string, error) {
	var decrypted []rune

	for _, letter := range encrypted {
		decryptedletter, err := modular.Exponentiation(int64(letter), privateExponent, modulus)

		if err != nil {
			return "", ErrorFailedToDecrypt
		}

		decrypted = append(decrypted, rune(decryptedletter))
	}

	return string(decrypted), nil
}
