package letter_combinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letterByDigit := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	if len(digits) == 1 {
		return letterByDigit[rune(digits[0])]
	}
	combinations := combinationsOf(letterByDigit[rune(digits[0])], letterByDigit[rune(digits[1])])
	for index := 2; index < len(digits); index++ {
		combinations = combinationsOf(combinations, letterByDigit[rune(digits[index])])
	}
	return combinations
}

func combinationsOf(first []string, other []string) []string {
	var combinations []string
	for _, firstCh := range first {
		for _, secondCh := range other {
			combinations = append(combinations, firstCh+secondCh)
		}
	}
	return combinations
}
