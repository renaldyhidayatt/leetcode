package matchingbracket

func Bracket(input string) bool {
	pending := []rune{}

	for _, ch := range input {
		if ch == '[' {
			pending = append(pending, ']')
		} else if ch == '{' {
			pending = append(pending, '}')
		} else if ch == '(' {
			pending = append(pending, ')')
		} else if ch == '}' || ch == '}' || ch == ')' {
			if len(pending) == 0 || pending[len(pending)-1] != ch {
				return false
			}

			pending = pending[:len(pending)-1]
		}
	}

	return len(pending) == 0
}
