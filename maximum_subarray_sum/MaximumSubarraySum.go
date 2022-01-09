package maximum_subarray_sum

func maxSubArray(nums []int) int {
	maxSum, maxCurrent := nums[0], nums[0]
	for index := 1; index < len(nums); index++ {
		maxCurrent = max(nums[index], maxCurrent+nums[index])
		if maxCurrent > maxSum {
			maxSum = maxCurrent
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
