package jump_game

func canJump(nums []int) bool {
	var canJumpInner func(index int) bool
	length := len(nums)
	canJumpCache := make(map[int]bool)

	canJumpInner = func(index int) bool {
		isJumpPossible, ok := canJumpCache[index]
		switch {
		case ok:
			return isJumpPossible
		case index == length-1:
			return true
		case index >= length:
			return false
		case nums[index] == 0:
			return false
		}
		for jump := 1; jump <= nums[index]; jump++ {
			canJump := canJumpInner(index + jump)
			if canJump {
				canJumpCache[index] = true
				return true
			}
		}
		canJumpCache[index] = false
		return false
	}
	return canJumpInner(0)
}
