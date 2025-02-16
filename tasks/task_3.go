package tasks

func IsValidBracketString(s string) bool {
	// Stack to store the opening brackets
	var (
		stack        []rune
		bracketPairs = map[rune]rune{
			'>': '<',
			'}': '{',
			']': '[',
		}
	)

	// Iterate over each character in the string
	for _, char := range s {
		// If it's an opening bracket, push it onto the stack
		if char == '<' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == '>' || char == '}' || char == ']' {
			// If it's a closing bracket, check for the matching opening bracket
			if len(stack) == 0 {
				return false // No opening bracket to match
			}

			// Pop from stack and check if it matches the corresponding opening bracket
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop the top element

			if top != bracketPairs[char] {
				return false
			}

			// Additional check to prevent curly brackets inside square brackets
			if char == ']' && top == '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
