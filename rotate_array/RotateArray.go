package rotate_array

func rotate(nums []int, k int) {
	for count := 1; count <= k; count++ {
		rotateRightBy1(nums)
	}
}

func rotateRightBy1(nums []int) {
	last := nums[len(nums)-1]
	for index := len(nums) - 1; index > 0; index-- {
		nums[index] = nums[index-1]
	}
	nums[0] = last
}
