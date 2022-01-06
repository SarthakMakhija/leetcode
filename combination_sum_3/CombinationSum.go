package combination_sum_3

func combinationSum3(k int, n int) [][]int {
	const maxNumberToUse = 9
	var combinationSum3Inner func(numbersUsed, remainingSum, current int) ([][]int, bool)

	combinationSum3Inner = func(numbersUsed, remainingSum, current int) ([][]int, bool) {
		switch {
		case numbersUsed > k:
			return [][]int{}, false
		case remainingSum == 0 && numbersUsed == k:
			return [][]int{}, true
		case remainingSum < 0:
			return [][]int{}, false
		}
		var combinations [][]int
		var possible bool
		for iterator := current; iterator <= maxNumberToUse; iterator++ {
			localCombinations, ok := combinationSum3Inner(numbersUsed+1, remainingSum-iterator, iterator+1)
			if ok {
				combinations, possible = merge(combinations, addTo(localCombinations, iterator)), ok
			}
		}
		return combinations, possible
	}
	combinations, _ := combinationSum3Inner(0, n, 1)
	return combinations
}

func merge(first [][]int, second [][]int) [][]int {
	merged := copy2D(first)
	for index := 0; index < len(second); index++ {
		merged = append(merged, second[index])
	}
	return merged
}

func copy2D(elements [][]int) [][]int {
	copied := make([][]int, len(elements))
	for index := range elements {
		copied[index] = make([]int, len(elements[index]))
		copy(copied[index], elements[index])
	}
	return copied
}

func addTo(elements [][]int, element int) [][]int {
	if len(elements) == 0 {
		return append(elements, []int{element})
	}
	copied := copy2D(elements)
	for index := 0; index < len(copied); index++ {
		replacement := make([]int, len(copied[index])+1)
		replacement[0] = element
		copy(replacement[1:], copied[index])
		copied[index] = replacement
	}
	return copied
}
