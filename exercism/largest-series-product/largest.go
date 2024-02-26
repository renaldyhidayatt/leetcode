package largestseriesproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("incorrect span")
	}

	d := make([]byte, len(digits))

	for i, v := range digits {
		if v < '0' || v > '9' {
			return 0, errors.New("not a digit")
		}

		d[i] = byte(v - '0')
	}

	max := int64(0)

	for i := 0; i < len(d)-span+1; i++ {
		v := int64(1)

		for n := i; n < i+span; n++ {
			v = v * int64(d[n])
		}

		if v > max {
			max = v
		}
	}
	return max, nil
}
