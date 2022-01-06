package evaluate_bracket_pairs

func evaluate(s string, knowledge [][]string) string {
	if len(s) == 0 {
		return ""
	}

	result, str, openingParentheses, closingParentheses, valueByPlaceholder := "", "", int32(40), int32(41), buildKnowledge(knowledge)
	for _, ch := range s {
		if ch == openingParentheses {
			result = result + str
			str = ""
			continue
		} else if ch == closingParentheses {
			v, ok := valueByPlaceholder[str]
			if ok {
				result = result + v
			} else {
				result = result + "?"
			}
			str = ""
			continue
		}
		str = str + string(ch)
	}
	return result + str
}

func buildKnowledge(knowledge [][]string) map[string]string {
	valueByPlaceholder := make(map[string]string)
	for _, kv := range knowledge {
		valueByPlaceholder[kv[0]] = kv[1]
	}
	return valueByPlaceholder
}
