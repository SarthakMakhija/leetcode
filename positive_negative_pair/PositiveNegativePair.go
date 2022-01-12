package positive_negative_pair

func findPairs(arr []int) []int {
	pairExistence := make(map[int]bool)
	var result []int

	for _, value := range arr {
		if value < 0 {
			if _, ok := pairExistence[0-value]; ok {
				result = append(result, value, 0-value)
				delete(pairExistence, 0-value)
			} else {
				pairExistence[value] = true
			}
		} else if value > 0 {
			if _, ok := pairExistence[0-value]; ok {
				result = append(result, 0-value, value)
				delete(pairExistence, 0-value)
			} else {
				pairExistence[value] = true
			}
		}
	}
	return result
}
