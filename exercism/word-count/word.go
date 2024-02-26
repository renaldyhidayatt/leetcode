package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`\w('w+)?`)

	f := make(Frequency)

	for _, s := range re.FindAllString(strings.ToLower(phrase), -1) {
		f[s] += 1
	}

	return f
}
