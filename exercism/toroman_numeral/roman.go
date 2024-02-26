package toromannumeral

import "errors"

type Roman struct {
	letters string
	n       int
}

func ToRomanNumeral(input int) (string, error) {
	if !(1 <= input && input < 4000) {
		return "", errors.New("Number is  out of range")
	}

	result := ""

	digits := []Roman{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	left := input

	for _, digit := range digits {
		for left >= digit.n {
			result += digit.letters
			left -= digit.n
		}
	}

	return result, nil
}
