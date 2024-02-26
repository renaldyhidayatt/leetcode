package strings

import (
	"errors"
	"regexp"
	"strings"
)

type IsogramOrder int

const (
	First IsogramOrder = iota + 1
	Second
	Third
)

func hasDigits(text string) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(text)
}

func hasSymbol(text string) bool {
	re := regexp.MustCompile(`[-!@#$%^&*()+]`)
	return re.MatchString(text)
}

func IsIsogram(text string, order IsogramOrder) (bool, error) {
	if order < First || order > Third {
		return false, errors.New("invalid isogram order provided")
	}

	text = strings.ToLower(text)
	text = strings.Join(strings.Fields(text), "")

	if hasDigits(text) || hasSymbol(text) {
		return false, errors.New("Cannot contain numbers or symbols")
	}

	letters := make(map[string]int)

	for _, c := range text {
		l := string(c)

		if _, ok := letters[l]; ok {
			letters[l] += 1

			if letters[l] > 3 {
				return false, nil
			}

			continue
		}

		letters[l] = 1
	}

	mapVals := make(map[int]bool)
	for _, v := range letters {
		mapVals[v] = true
	}

	if _, ok := mapVals[int(order)]; ok && len(mapVals) == 1 {
		return true, nil
	}

	return false, nil
}
