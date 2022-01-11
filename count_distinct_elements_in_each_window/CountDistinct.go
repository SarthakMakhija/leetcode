package count_distinct_elements_in_each_window

func countDistinct(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var distinct []int
	frequency, totalWindows := frequencyInWindow1(k, nums), len(nums)-k+1
	distinct = append(distinct, len(frequency))

	for window := 2; window <= totalWindows; window++ {
		startIndex, previousWindowStartIndex := window-1, window-1-1
		endIndex := startIndex + k - 1

		frequency[nums[endIndex]] = frequency[nums[endIndex]] + 1
		frequency[nums[previousWindowStartIndex]] = frequency[nums[previousWindowStartIndex]] - 1

		if frequency[nums[previousWindowStartIndex]] == 0 {
			delete(frequency, nums[previousWindowStartIndex])
		}
		distinct = append(distinct, len(frequency))
	}

	return distinct
}

func frequencyInWindow1(windowSize int, nums []int) map[int]int {
	frequency := make(map[int]int)
	for index := 0; index < windowSize; index++ {
		frequency[nums[index]] = frequency[nums[index]] + 1
	}
	return frequency
}
