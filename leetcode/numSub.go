package main

const mod = 1000000007

func numSub(s string) int {
	len := 0
	result := int64(0)
	for _, c := range s {
		if c == '1' {
			len++
			continue
		}
		if len > 0 {
			result += int64(len*(len+1)) / 2
			len = 0
		}
	}
	result += int64(len*(len+1)) / 2
	return int(result % mod)
}
