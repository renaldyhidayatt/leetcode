package validparentheses

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := []rune{}
	for _, t := range s {
		if t == ')' {
			if len(stack) == 0 {
				return false
			}
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if current != '(' {
				return false
			}
		} else if t == '}' {
			if len(stack) == 0 {
				return false
			}
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if current != '{' {
				return false
			}
		} else if t == ']' {
			if len(stack) == 0 {
				return false
			}
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if current != '[' {
				return false
			}
		} else {
			stack = append(stack, t)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func RunningValidParentheses() {
	fmt.Println(isValid("()"))
}
