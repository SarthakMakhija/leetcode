package house_robber

func rob(nums []int) int {
	var robInner func(currentIndex int) int

	length := len(nums)
	robCache := make(map[int]int)
	robInner = func(currentIndex int) int {
		robAmount, ok := robCache[currentIndex]
		switch {
		case ok:
			return robAmount
		case currentIndex == length-1:
			return nums[currentIndex]
		case currentIndex >= length:
			return 0
		}
		amount := -1
		for iterator := currentIndex; iterator < length; iterator++ {
			robbed := nums[iterator] + robInner(iterator+2)
			if robbed > amount {
				amount = robbed
			}
		}
		robCache[currentIndex] = amount
		return robCache[currentIndex]
	}
	return robInner(0)
}
