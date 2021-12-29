package find_peak

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearchIn(nums)
}

func binarySearchIn(nums []int) int {
	pivotIndex := len(nums) / 2
	if len(nums) == 1 {
		return 0
	} else if nums[pivotIndex] < nums[pivotIndex-1] {
		return binarySearchIn(nums[0:pivotIndex])
	} else if pivotIndex == len(nums)-1 {
		if nums[pivotIndex] > nums[pivotIndex-1] {
			return pivotIndex
		}
		return -1
	} else if nums[pivotIndex] < nums[pivotIndex+1] {
		nextIndex := pivotIndex + 1
		index := binarySearchIn(nums[nextIndex:])
		if index != -1 {
			return nextIndex + index
		}
		return index
	} else if nums[pivotIndex] > nums[pivotIndex-1] && nums[pivotIndex] > nums[pivotIndex+1] {
		return pivotIndex
	}
	return -1
}
