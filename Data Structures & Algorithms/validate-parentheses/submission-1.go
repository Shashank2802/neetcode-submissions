func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
