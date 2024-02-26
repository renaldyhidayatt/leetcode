package runlength

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	var encoded string
	var current rune
	var count int

	for _, r := range input {
		if r != current {
			if count > 0 {
				encoded += encode(current, count)
			}
			current = r
			count = 1
		} else {
			count++
		}
	}
	encoded += encode(current, count)

	return encoded
}

func encode(r rune, count int) string {
	if count == 1 {
		return string(r)
	} else {
		return fmt.Sprintf("%d%c", count, r)
	}
}

func RunLengthDecode(input string) string {
	var decoded string

	countStr := ""

	var count int

	for _, r := range input {
		if unicode.IsDigit(r) {
			countStr += string(r)
		} else {
			if countStr == "" {
				count = 1
			} else {
				count, _ = strconv.Atoi(countStr)
			}

			decoded += decode(r, count)
			countStr = ""
		}
	}
	return decoded
}

func decode(r rune, count int) string {
	return strings.Repeat(string(r), count)
}
