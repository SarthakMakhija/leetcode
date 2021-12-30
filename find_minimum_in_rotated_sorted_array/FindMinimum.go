package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if !isRotated(nums) {
		return nums[0]
	}
	return binarySearchMinIn(nums, nums[0])
}

func isRotated(nums []int) bool {
	if nums[0] < nums[len(nums)-1] {
		return false
	}
	return true
}

func binarySearchMinIn(nums []int, firstIndexValue int) int {
	for len(nums) > 0 {
		pivotIndex := len(nums) / 2
		if nums[pivotIndex] > firstIndexValue {
			nums = nums[pivotIndex+1:]
		} else {
			if len(nums) == 1 {
				return nums[0]
			}
			if nums[pivotIndex-1] < nums[pivotIndex] {
				nums = nums[0:pivotIndex]
			} else {
				return nums[pivotIndex]
			}
		}
	}
	return -1
}
