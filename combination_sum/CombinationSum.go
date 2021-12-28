package combination_sum

type combinationPossibility struct {
	possible    bool
	combination [][]int
}

type frequency struct {
	countByValue map[int]int
}

func combinationSum(candidates []int, target int) [][]int {
	var combinationSumInner func(remaining int) combinationPossibility

	combinationSumCache := make(map[int]combinationPossibility)
	combinationSumInner = func(remaining int) combinationPossibility {
		cachedCombination, ok := combinationSumCache[remaining]
		switch {
		case ok:
			return cachedCombination
		case remaining == 0:
			return combinationPossibility{possible: true, combination: [][]int{}}
		case remaining < 0:
			return combinationPossibility{possible: false, combination: [][]int{}}
		}
		var combinations [][]int
		var possible bool

		for _, value := range candidates {
			combinationPossibility := combinationSumInner(remaining - value)
			localCombination, ok := combinationPossibility.combination, combinationPossibility.possible
			if ok {
				combinations = merge(combinations, addTo(localCombination, value))
				possible = true
			}
		}
		combinationSumCache[remaining] = combinationPossibility{possible: possible, combination: combinations}
		return combinationSumCache[remaining]
	}
	return unique(combinationSumInner(target).combination)
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

func unique(combinations [][]int) [][]int {
	var frequencies []frequency
	var uniqueCombinations [][]int

	for _, combination := range combinations {
		frequency := frequencyOf(combination)
		if !contains(frequencies, frequency) {
			frequencies = append(frequencies, frequency)
			uniqueCombinations = append(uniqueCombinations, combination)
		}
	}
	return uniqueCombinations
}

func contains(frequencies []frequency, frequency frequency) bool {
	for _, f := range frequencies {
		if matches(f, frequency) {
			return true
		}
	}
	return false
}

func matches(one frequency, other frequency) bool {
	if len(one.countByValue) != len(other.countByValue) {
		return false
	}
	for k, value := range one.countByValue {
		if other.countByValue[k] != value {
			return false
		}
	}
	return true
}

func frequencyOf(values []int) frequency {
	countByValue := make(map[int]int)
	for _, value := range values {
		countByValue[value] = countByValue[value] + 1
	}
	return frequency{countByValue: countByValue}
}
