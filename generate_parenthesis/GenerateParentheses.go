package generate_parenthesis

func generateParenthesis(n int) []string {
	var parenthesis []string

	var generateParenthesisInner func(remainingOpeningParenthesis, remainingClosingParenthesis int, generated string)
	generateParenthesisInner = func(remainingOpeningParenthesis, remainingClosingParenthesis int, generated string) {
		if remainingOpeningParenthesis == 0 && remainingClosingParenthesis == 0 {
			if isValidParenthesis(generated) {
				parenthesis = append(parenthesis, generated)
			}
		}
		if remainingOpeningParenthesis > 0 {
			generateParenthesisInner(
				remainingOpeningParenthesis-1,
				remainingClosingParenthesis,
				generated+"(",
			)
		}
		if remainingClosingParenthesis > 0 && len(generated) != 0 {
			generateParenthesisInner(
				remainingOpeningParenthesis,
				remainingClosingParenthesis-1,
				generated+")",
			)
		}
	}
	generateParenthesisInner(n, n, "")
	return parenthesis
}

func isValidParenthesis(parenthesis string) bool {
	var stack []rune
	for _, ch := range parenthesis {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
