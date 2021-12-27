package first_and_last_position

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return binarySearchIn(nums, target)
}

func binarySearchIn(nums []int, target int) []int {
	pivotIndex := len(nums) / 2
	switch {
	case len(nums) == 0:
		return []int{-1, -1}
	case target < nums[pivotIndex]:
		return binarySearchIn(nums[0:pivotIndex], target)
	case target > nums[pivotIndex]:
		nextIndex := pivotIndex + 1
		return addOffsetTo(binarySearchIn(nums[pivotIndex+1:], target), nextIndex)
	case target == nums[pivotIndex]:
		backIndex := moveBackIn(pivotIndex, func(index int) bool {
			return index >= 0 && nums[index] == target
		})
		frontIndex := moveFrontIn(pivotIndex, func(index int) bool {
			return index < len(nums) && nums[index] == target
		})
		return []int{backIndex, frontIndex}
	}
	return []int{-1, -1}
}

func addOffsetTo(indices []int, offsetIndex int) []int {
	if indices[0] != -1 && indices[1] != -1 {
		withOffset := make([]int, len(indices))
		copy(withOffset, indices)

		withOffset[0] = withOffset[0] + offsetIndex
		withOffset[1] = withOffset[1] + offsetIndex
		return withOffset
	}
	return indices
}

func moveBackIn(fromIndex int, condition func(index int) bool) int {
	backIndex := fromIndex
	for ; condition(backIndex); backIndex-- {
	}
	backIndex = backIndex + 1
	return backIndex
}

func moveFrontIn(fromIndex int, condition func(index int) bool) int {
	frontIndex := fromIndex
	for ; condition(frontIndex); frontIndex++ {
	}
	frontIndex = frontIndex - 1
	return frontIndex
}
