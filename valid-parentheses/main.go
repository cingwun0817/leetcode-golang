// FILO  (First In Last Out)
func isValid(s string) bool {
	stack := []int32{}

	for _, c := range s {
		if c != '(' && c != ')' && c != '{' && c != '}' && c != '[' && c != ']' {
			continue
		}

		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			//fmt.Println(stack)
			if len(stack) == 0 {
				return false
			}

			if c == ')' && stack[len(stack)-1] != '(' {
				return false
			}

			if c == '}' && stack[len(stack)-1] != '{' {
				return false
			}

			if c == ']' && stack[len(stack)-1] != '[' {
				return false
			}

			stack = stack[0 : len(stack)-1]

			//fmt.Println(stack)
		}
	}

	return len(stack) == 0
}