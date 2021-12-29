package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length, front, tail := len(nums), 0, 0
	for front < length {
		tail = move(tail, func(tailIndex int) bool {
			return tailIndex < length && nums[tailIndex] == nums[front]
		})
		count := tail - front
		if count <= 2 {
			front = tail
		} else {
			localCount := 0
			front = move(front, func(frontIndex int) bool {
				localCount = localCount + 1
				return localCount <= 2
			})
			copy(nums[front:], nums[tail:])
			length = length - (tail - front)
			tail = front
		}
	}
	return length
}

func move(index int, f func(index int) bool) int {
	iteratingIndex := index
	for f(iteratingIndex) {
		iteratingIndex = iteratingIndex + 1
	}
	return iteratingIndex
}
