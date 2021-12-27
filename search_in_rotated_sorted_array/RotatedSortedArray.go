package search_in_rotated_sorted_array

import (
	"sort"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if !isRotated(nums) {
		return binarySearchIn(nums, target, 0)
	}
	return searchInRotated(nums, target)
}

func isRotated(nums []int) bool {
	return nums[0] > nums[len(nums)-1]
}

func binarySearchIn(nums []int, target int, addIndex int) int {
	targetIndex := sort.Search(len(nums), func(index int) bool {
		return nums[index] >= target
	})
	if targetIndex < len(nums) && nums[targetIndex] == target {
		return targetIndex + addIndex
	}
	return -1
}

func searchInRotated(nums []int, target int) int {
	pivotIndex, initialIndex := len(nums)/2, 0

	for pivotIndex < len(nums) {
		if nums[pivotIndex] > nums[0] && nums[pivotIndex] > nums[len(nums)-1] {
			if target >= nums[initialIndex] && target <= nums[pivotIndex] {
				return binarySearchIn(nums[initialIndex:pivotIndex+1], target, initialIndex)
			}
			nextIndex := pivotIndex + 1
			pivotIndex = nextIndex + len(nums[nextIndex:])/2
			initialIndex = nextIndex
		} else {
			backIndex := moveBackIn(pivotIndex, func(index int) bool {
				return nums[index] < nums[0]
			})
			if target >= nums[backIndex] && target <= nums[len(nums)-1] {
				return binarySearchIn(nums[backIndex:], target, backIndex)
			} else {
				return binarySearchIn(nums[initialIndex:backIndex], target, initialIndex)
			}
		}
	}
	return -1
}

func moveBackIn(fromIndex int, condition func(index int) bool) int {
	backIndex := fromIndex
	for ; condition(backIndex); backIndex-- {
	}
	backIndex = backIndex + 1
	return backIndex
}
