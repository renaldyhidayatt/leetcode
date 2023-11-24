package on299

import "container/list"

func Calculate(s string) int {
	i, stack, result, sign := 0, list.New(), 0, 1

	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else if s[i] <= '9' && s[i] >= '0' {
			base, v := 10, int(s[i]-'0')

			for i+1 < len(s) && s[i+1] <= '9' && s[i+1] >= '0' {
				v = v*base + int(s[i+1]-'0')
				i++
			}

			result += v * sign
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '(' {
			stack.PushBack(result)
			stack.PushBack(sign)
			result = 0
			sign = 1
			i++
		} else if s[i] == ')' {
			result = result*stack.Remove(stack.Back()).(int) + stack.Remove(stack.Back()).(int)
			i++
		}
	}

	return result
}
