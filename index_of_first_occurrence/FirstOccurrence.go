package index_of_first_occurrence

func firstOccurrence(elements []int, x int) int {
	if len(elements) == 0 {
		return -1
	}
	low := 0
	high := len(elements) - 1
	return index(elements, x, low, high)
}

func index(elements []int, x int, low int, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if x > elements[mid] {
			low = mid + 1
		} else if x < elements[mid] {
			high = mid - 1
		} else {
			if mid == 0 || elements[mid-1] != elements[mid] {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}
