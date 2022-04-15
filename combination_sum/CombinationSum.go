package combination_sum

import (
	"reflect"
)

func combinationSum(candidates []int, target int) [][]int {
	var empty [][]int

	var combinationSumInner func(remainingSum int) ([][]int, bool)
	combinationSumInner = func(remainingSum int) ([][]int, bool) {
		switch {
		case remainingSum == 0:
			return empty, true
		}
		var allCombinations [][]int
		for _, element := range candidates {
			if remainingSum-element >= 0 {
				combinations, possible := combinationSumInner(remainingSum - element)
				if possible {
					allCombinations = merge(allCombinations, add(combinations, element))
				}
			}
		}
		if len(allCombinations) == 0 {
			return allCombinations, false
		}
		return allCombinations, true
	}
	allCombinations, _ := combinationSumInner(target)
	return allCombinations
}

func add(allCombinations [][]int, element int) [][]int {
	if len(allCombinations) == 0 {
		return append(allCombinations, []int{element})
	}
	var newCombinations [][]int
	for index := 0; index < len(allCombinations); index++ {
		newCombinations = append(newCombinations, append(allCombinations[index], element))
	}
	return newCombinations
}

func merge(one [][]int, other [][]int) [][]int {
	var merged [][]int
	var frequencies []map[int]int

	for index := 0; index < len(one); index++ {
		if f := frequencyOf(one[index]); !containsFrequency(f, frequencies) {
			merged = append(merged, one[index])
			frequencies = append(frequencies, f)
		}
	}
	for index := 0; index < len(other); index++ {
		if f := frequencyOf(other[index]); !containsFrequency(f, frequencies) {
			merged = append(merged, other[index])
			frequencies = append(frequencies, f)
		}
	}
	return merged
}

func frequencyOf(elements []int) map[int]int {
	var frequency = make(map[int]int)
	for _, element := range elements {
		if count, ok := frequency[element]; ok {
			frequency[element] = count + 1
		} else {
			frequency[element] = 1
		}
	}
	return frequency
}

func containsFrequency(frequency map[int]int, frequencies []map[int]int) bool {
	for _, existingFrequency := range frequencies {
		if reflect.DeepEqual(existingFrequency, frequency) {
			return true
		}
	}
	return false
}
