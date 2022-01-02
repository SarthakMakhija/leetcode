package sort_colors

func sortColors(nums []int) {
	zeroCount, oneCount, twoCount := colorCount(nums)
	for index := 0; index < len(nums); {
		index = fill(nums, zeroCount, 0, index)
		index = fill(nums, oneCount, 1, index)
		index = fill(nums, twoCount, 2, index)
	}
}

func fill(nums []int, totalCount int, value int, fromIndex int) int {
	index := fromIndex
	for count := 0; count < totalCount; count++ {
		nums[index] = value
		index = index + 1
	}
	return index
}

func colorCount(nums []int) (int, int, int) {
	zeroCount, oneCount, twoCount := 0, 0, 0
	for _, value := range nums {
		if value == 0 {
			zeroCount = zeroCount + 1
		} else if value == 1 {
			oneCount = oneCount + 1
		} else if value == 2 {
			twoCount = twoCount + 1
		}
	}
	return zeroCount, oneCount, twoCount
}
