package main

func checkZeroOnes(s string) bool {
	zeros := 0
	ones := 0
	i := 0
	for i < len(s) {
		if s[i] == '0' {
			cnt := 0
			for ; i < len(s) && s[i] == '0'; i++ {
				cnt++
			}

			if cnt > zeros {
				zeros = cnt
			}
		} else {
			cnt := 0
			for ; i < len(s) && s[i] == '1'; i++ {
				cnt++
			}
			if cnt > ones {
				ones = cnt
			}
		}
	}

	return ones > zeros
}
