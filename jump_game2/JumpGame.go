package jump_game2

import (
	"math"
)

func jump(nums []int) int {
	var jumpInner func(index int) int
	jumpNotPossible := math.MaxInt32

	length := len(nums)
	jumpCache := make(map[int]int)

	jumpInner = func(index int) int {
		cachedJumps, ok := jumpCache[index]
		switch {
		case ok:
			return cachedJumps
		case index == length-1:
			return 0
		case index >= length:
			return jumpNotPossible
		case nums[index] == 0:
			return jumpNotPossible
		}
		minimumJumps := -1
		for jump := 1; jump <= nums[index]; jump++ {
			jumps := jumpInner(index+jump) + 1
			if minimumJumps == -1 || jumps < minimumJumps {
				minimumJumps = jumps
			}
		}
		jumpCache[index] = minimumJumps
		return jumpCache[index]
	}
	return jumpInner(0)
}
