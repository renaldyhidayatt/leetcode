package parallelletterfrequency

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}

	for _, r := range s {
		m[r]++
	}

	return m
}

func ConccurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}

	c := make(chan FreqMap, 10)

	for _, s := range ss {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for range ss {
		for k, v := range <-c {
			m[k] += v
		}
	}

	return m
}
