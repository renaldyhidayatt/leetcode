package conversion

import "errors"

var (
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"} // 1 - 9
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} // 10 - 90
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} // 100 - 900
	r3 = []string{"", "M", "MM", "MMM"}                                       // 1,000 - 3,000
)

func IntToRoman(n int) (string, error) {

	if n < 1 || n > 3999 {
		return "", errors.New("integer must have value between 1 and 3999")
	}

	return r3[n%1e4/1e3] + r2[n%1e3/1e2] + r1[n%100/10] + r0[n%10], nil
}
