package validParentheses

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
		} else if v == ']' || v == ')' || v == '}' {

			if len(stack) == 0 {
				return false
			}

			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if v == ']' && pop != '[' {
				return false
			}

			if v == ')' && pop != '(' {
				return false
			}

			if v == '}' && pop != '{' {
				return false
			}

		}

	}

	return len(stack) == 0
}
