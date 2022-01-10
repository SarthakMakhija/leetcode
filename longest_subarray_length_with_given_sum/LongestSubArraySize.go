package longest_subarray_length_with_given_sum

func longestSubArrayLengthWithSum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
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
