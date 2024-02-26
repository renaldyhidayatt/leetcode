package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	isSilence := trimmedRemark == ""
	hasLetters := strings.IndexFunc(trimmedRemark, unicode.IsLetter) >= 0
	isYelling := hasLetters && strings.ToUpper(trimmedRemark) == trimmedRemark
	isQuestion := strings.HasSuffix(trimmedRemark, "?")

	switch {
	case isSilence:
		return "Fine. Be that way!"
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
