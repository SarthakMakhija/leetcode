package sort_an_array_according_to_other

import (
	"sort"
)

func sortA1ByA2(A1 []int, A2 []int) []int {
	relativeSorted := make([]int, len(A1))
	frequency := make(map[int]int)

	for _, value := range A1 {
		frequency[value] = frequency[value] + 1
	}
	index := 0
	for _, value := range A2 {
		if countOfValue, ok := frequency[value]; ok {
			index = fill(relativeSorted, value, countOfValue, index)
			delete(frequency, value)
		}
	}
	keys := sortByKey(frequency)
	for _, key := range keys {
		index = fill(relativeSorted, key, frequency[key], index)
	}

	return relativeSorted
}

func sortByKey(frequency map[int]int) []int {
	sortedKeyIndex, remaining := 0, make([]int, len(frequency))
	for key := range frequency {
		remaining[sortedKeyIndex] = key
		sortedKeyIndex = sortedKeyIndex + 1
	}
	sort.Ints(remaining)
	return remaining
}

func fill(nums []int, value int, countOfValue int, fromIndex int) int {
	index := fromIndex
	for count := 1; count <= countOfValue; count++ {
		nums[index] = value
		index = index + 1
	}
	return index
}
