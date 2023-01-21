package lettercombinations

import "fmt"

func letterCombinations(digits string) []string {
	phone := [][]rune{
		{' ', '\x00', '\x00', '\x00'},    //0
		{'\x00', '\x00', '\x00', '\x00'}, //1
		{'a', 'b', 'c', '\x00'},          //2
		{'d', 'e', 'f', '\x00'},          //3
		{'g', 'h', 'i', '\x00'},          //4
		{'j', 'k', 'l', '\x00'},          //5
		{'m', 'n', 'o', '\x00'},          //6
		{'p', 'q', 'r', 's'},             //7
		{'t', 'u', 'v', '\x00'},          //8
		{'w', 'x', 'y', 'z'},             //9
	}

	result := make([]string, 0)
	if len(digits) <= 0 {
		result = append(result, "")
		return result
	}

	for i := 0; i < len(digits); i++ {
		if digits[i] < '0' || digits[i] > '9' {
			return make([]string, 0)
		}
		d := digits[i] - '0'
		if len(result) <= 0 {
			for j := 0; j < 4 && phone[d][j] != '\x00'; j++ {
				s := string(phone[d][j])
				result = append(result, s)
			}
			continue
		}
		r := make([]string, 0)
		for j := 0; j < len(result); j++ {
			for k := 0; k < 4 && phone[d][k] != '\x00'; k++ {
				s := result[j] + string(phone[d][k])
				r = append(r, s)
			}
			result = r
		}
	}
	return result

}

func RunningLetterCombinations() {
	s := "23"
	fmt.Println(letterCombinations(s))
}
