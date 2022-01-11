package longest_consecutive_subsequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	occurrence := make(map[int]bool)
	for _, num := range nums {
		if _, ok := occurrence[num]; !ok {
			occurrence[num] = true
		}
	}

	longestSubsequence := 0
	for num := range occurrence {
		length := 1
		if _, ok := occurrence[num-1]; !ok {
			for next := num + 1; occurrence[next]; next++ {
				length = length + 1
			}
		}
		if length > longestSubsequence {
			longestSubsequence = length
		}
	}
	return longestSubsequence
}
