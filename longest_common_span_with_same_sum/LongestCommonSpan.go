package longest_common_span_with_same_sum

func longestCommonSum(arr1 []bool, arr2 []bool) int {
	diff := make([]int, len(arr1))
	for index := 0; index < len(arr1); index++ {
		switch {
		case arr1[index] && arr2[index]:
			diff[index] = 0
		case !arr1[index] && !arr2[index]:
			diff[index] = 0
		case arr1[index] && !arr2[index]:
			diff[index] = 1
		case !arr1[index] && arr2[index]:
			diff[index] = -1
		}
	}
	return longestSubArrayLengthWithSumZero(diff)
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
