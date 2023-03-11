package lettercasepermutation

import "strings"

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}
	res, pos, c := []string{}, []int{}, []int{}
	SS := strings.ToLower(S)
	for i := 0; i < len(SS); i++ {
		if isLowerLetter(SS[i]) {
			pos = append(pos, i)
		}
	}
	for i := 0; i <= len(pos); i++ {
		findLetterCasePermutation(SS, pos, i, 0, c, &res)
	}
	return res
}

func findLetterCasePermutation(s string, pos []int, target, index int, c []int, res *[]string) {
	if len(c) == target {
		b := []byte(s)
		for _, v := range c {
			b[pos[v]] -= 'a' - 'A'
		}
		*res = append(*res, string(b))
		return
	}
	for i := index; i < len(pos)-(target-len(c))+1; i++ {
		c = append(c, i)
		findLetterCasePermutation(s, pos, target, i+1, c, res)
		c = c[:len(c)-1]
	}
}

func isLowerLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
