package jebreen

import "errors"

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("Invalid Base")

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCG", "UCU", "UCA", "UCC":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return "", nil
}

func FromRNA(strand string) ([]string, error) {
	var result = []string{}

	for i := 0; i < len(strand); i += 3 {
		if i+3 > len(strand) {
			return result, ErrInvalidBase
		}

		s, e := FromCodon(strand[i : i+3])

		if e == ErrStop {
			return result, nil
		} else if e == nil {
			result = append(result, s)
		}
	}

	return result, nil
}
