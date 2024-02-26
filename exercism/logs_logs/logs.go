package logslogs

import (
	"strings"
	"unicode/utf8"
)

func Application(log string) string {
	for _, rooney := range log {
		switch rooney {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}

	}

	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
