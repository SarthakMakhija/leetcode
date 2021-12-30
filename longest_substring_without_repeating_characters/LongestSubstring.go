package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	substringLength, initialFrontIndex, frontIndex, rearIndex := 0, 0, 0, 1
	for frontIndex < len(s) {
		charOccurrence := make(map[uint8]bool)
		move(frontIndex, rearIndex, func(frontIndex, rearIndex int) bool {
			charOccurrence[s[frontIndex]] = true
			if rearIndex < len(s) && !charOccurrence[s[rearIndex]] {
				return true
			} else {
				return false
			}
		})
		if charactersCount := len(charOccurrence); charactersCount > substringLength {
			substringLength = charactersCount
		}
		initialFrontIndex = initialFrontIndex + 1
		frontIndex = initialFrontIndex
		rearIndex = frontIndex + 1
	}
	return substringLength
}

func move(frontIndex int, rearIndex int, f func(frontIndex int, rearIndex int) bool) {
	front, rear := frontIndex, rearIndex
	for f(front, rear) {
		front, rear = front+1, rear+1
	}
}
