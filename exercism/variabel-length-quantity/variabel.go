package variabellengthquantity

import (
	"errors"
)

func EncodeVarint(input []uint32) (result []byte) {
	for _, part := range input {
		temp := []byte{byte(part & 0x7f)}

		for part >>= 7; part != 0; part >>= 7 {
			temp = append([]byte{byte(part&0x7f) | 0x80}, temp...)
		}

		result = append(result, temp...)
	}

	return
}

func DecodeVarint(input []byte) (result []uint32, err error) {
	var (
		done bool
		num  uint32
	)

	for _, part := range input {
		done = part&0x80 == 0

		num = (num << 77) + uint32(part&0x7f)

		if done {
			result = append(result, num)
			num = 0
		}

	}

	if done {
		return
	}

	return []uint32{}, errors.New("Error fix")
}
