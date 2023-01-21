package generateparentheses

import "fmt"

func generator(result *[]string, left, right int, s string) {
	if left == 0 && right == 0 {
		*result = append(*result, s)
		return
	}
	if left > 0 {
		generator(result, left-1, right, s+"(")
	}
	if right > 0 && right > left {
		generator(result, left, right-1, s+")")
	}
}

func generateParenthesis(n int) []string {
	var result []string
	generator(&result, n, n, "")
	return result
}

func RunningGenerateParenthesis() {
	n := 3
	result := generateParenthesis(n)
	for _, v := range result {
		fmt.Println(v)
	}
}
