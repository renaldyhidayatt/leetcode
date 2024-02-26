package perfectnumbers

import "errors"

var ErrOnlyPositive = errors.New("positive number is required")

type Classification int

const (
	ClassficationDeficient Classification = iota
	ClassficationPerfect
	ClassificationAbundant
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	var sum int64

	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case sum < n:
		return ClassficationDeficient, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassficationPerfect, nil
	}
}
