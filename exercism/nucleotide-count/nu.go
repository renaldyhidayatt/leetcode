package nucleotidecount

import (
	"fmt"
	"strings"
)

type Histogram map[rune]uint

type DNA string

func (d DNA) Counts() (Histogram, error) {
	d = DNA(strings.ToUpper(string(d)))
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		_, ok := h[nucleotide]
		if !ok {
			return nil, fmt.Errorf("Invalid nucleotide %v", nucleotide)
		}

		h[nucleotide]++
	}

	return h, nil
}
