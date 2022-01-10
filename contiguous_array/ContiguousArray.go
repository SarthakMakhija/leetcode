package contiguous_array

func findMaxLength(nums []int) int {
	backup := nums[:]
	for index, value := range backup {
		if value == 0 {
			backup[index] = -1
		}
	}
	return longestSubArrayLengthWithSumZero(nums)
}

func longestSubArrayLengthWithSumZero(nums []int) int {
	k := 0
	prefixSum := make(map[int]int) //subarray sum, index
	subArrayLength, currentSum := 0, 0

	for index, num := range nums {
		currentSum = currentSum + num
		if currentSum == k {
			if index+1 > subArrayLength {
				subArrayLength = index + 1
			}
		}
		if existingSumIndex, ok := prefixSum[currentSum-k]; ok {
			if index-existingSumIndex > subArrayLength {
				subArrayLength = index - existingSumIndex
			}
		}
		if _, ok := prefixSum[currentSum]; !ok {
			prefixSum[currentSum] = index
		}
	}
	return subArrayLength
}
