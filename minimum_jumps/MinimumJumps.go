package minimum_jumps

func jump(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	length := len(nums)
	jumpCache := make(map[int]int)

	var jumpInner func(currentIndex int) int
	jumpInner = func(currentIndex int) int {
		cached, ok := jumpCache[currentIndex]
		switch {
		case ok:
			return cached
		case currentIndex == length-1:
			return 0
		}
		minJumps := -1
		for jump := 1; jump <= nums[currentIndex]; jump++ {
			if currentIndex+jump < length {
				if jumps := jumpInner(currentIndex + jump); jumps != -1 {
					totalJumps := 1 + jumps
					if minJumps == -1 || totalJumps < minJumps {
						minJumps = totalJumps
					}
				}
			}
		}
		jumpCache[currentIndex] = minJumps
		return jumpCache[currentIndex]
	}
	return jumpInner(0)
}
