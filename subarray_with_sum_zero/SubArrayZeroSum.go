package subarray_with_sum_zero

func checkSubarraySumZero(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	prefixSum := make(map[int]bool)
	const expectedSum = 0
	currentSum := 0

	for _, num := range nums {
		if num == 0 {
			return true
		}
		currentSum = currentSum + num
		if currentSum == 0 {
			return true
		}
		if prefixSum[currentSum-expectedSum] {
			return true
		}
		prefixSum[currentSum] = true
	}
	return false
}
