package day9

func isValid(s string) bool {
	var stack []rune

	for _, v := range s {
		switch v {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			size := len(stack)
			if stack[size-1] != v || size == 0 {
				return false
			}
			stack = stack[:size-1]
		}
	}

	return len(stack) == 0
}
