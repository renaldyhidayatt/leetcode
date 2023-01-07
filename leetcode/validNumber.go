package main

import (
	"strings"
	"unicode"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	ls, pos := len(s), 0
	if ls == 0 {
		return false
	}
	if s[pos] == '+' || s[pos] == '-' {
		pos++
	}
	isNumeric := false
	for pos < ls && unicode.IsDigit(rune(s[pos])) {
		pos++
		isNumeric = true
	}
	if pos < ls && s[pos] == '.' {
		pos++
		for pos < ls && unicode.IsDigit(rune(s[pos])) {
			pos++
			isNumeric = true
		}
	} else if pos < ls && s[pos] == 'e' && isNumeric {
		isNumeric = false
		pos++
		if pos < ls && (s[pos] == '+' || s[pos] == '-') {
			pos++
		}
		for pos < ls && unicode.IsDigit(rune(s[pos])) {
			pos++
			isNumeric = true
		}
	}
	if pos == ls && isNumeric {
		return true
	}
	return false
}
