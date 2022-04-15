package subsets_2

import "reflect"

func subsetsWithDup(nums []int) [][]int {
	var subsetsInner func(remainingNums []int) [][]int
	subsetsInner = func(remainingNums []int) [][]int {
		if len(remainingNums) == 1 {
			return [][]int{{}, {remainingNums[0]}}
		}
		return addSubsetsTo(subsetsInner(remainingNums[1:]), remainingNums[0])
	}
	var uniqueSubsets [][]int
	var frequencies []map[int]int

	for _, subset := range subsetsInner(nums) {
		if f := frequencyOf(subset); !containsFrequency(f, frequencies) {
			uniqueSubsets = append(uniqueSubsets, subset)
			frequencies = append(frequencies, f)
		}
	}
	return uniqueSubsets
}

func addSubsetsTo(subsets [][]int, num int) [][]int {
	var allSubsets [][]int
	numberSubset := []int{num}
	for index := 0; index < len(subsets); index++ {
		allSubsets = append(allSubsets, subsets[index])
		if len(subsets[index]) != 0 {
			newSubset := append(numberSubset, subsets[index]...)
			allSubsets = append(allSubsets, newSubset)
		}
	}
	return append(allSubsets, numberSubset)
}

func frequencyOf(subset []int) map[int]int {
	var frequency = make(map[int]int)
	for _, element := range subset {
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
