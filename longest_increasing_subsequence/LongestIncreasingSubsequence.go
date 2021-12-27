package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	var lengthOfLISInner func(index int, value int) int

	lengthOfLISCache := make(map[int]int)
	lengthOfLISInner = func(index int, value int) int {
		max, ok := lengthOfLISCache[index]
		if ok {
			return max
		}
		values := allValuesGreaterThanFromIndex(nums, value, index+1)
		if len(values) == 0 {
			return 1
		}
		maxLength := -1
		for _, value := range values {
			length := lengthOfLISInner(value.index, value.value) + 1
			if length > maxLength {
				maxLength = length
			}
		}
		lengthOfLISCache[index] = maxLength
		return lengthOfLISCache[index]
	}
	maximumSubsequenceLength := -1
	for index, value := range nums {
		length := lengthOfLISInner(index, value)
		if length > maximumSubsequenceLength {
			maximumSubsequenceLength = length
		}
	}
	return maximumSubsequenceLength
}

type Value struct {
	index, value int
}

func allValuesGreaterThanFromIndex(nums []int, value int, index int) []Value {
	var values []Value
	for iterator := index; iterator < len(nums); iterator++ {
		if nums[iterator] > value {
			values = append(values, Value{index: iterator, value: nums[iterator]})
		}
	}
	return values
}
