package count_occurrence_in_sorted

func countOccurrence(elements []int, x int) int {
	if len(elements) == 0 {
		return -1
	}
	low := 0
	high := len(elements) - 1

	earliest, last := earliestAndLastOccurrence(elements, x, low, high)
	if earliest == -1 {
		return 0
	}
	first := firstOccurrence(elements, x, 0, earliest-1)
	if first == -1 {
		return last - earliest + 1
	}
	return last - first + 1
}

func earliestAndLastOccurrence(elements []int, x int, low int, high int) (int, int) {
	earliestOccurrence := -1
	for low <= high {
		mid := (low + high) / 2
		if x > elements[mid] {
			low = mid + 1
		} else if x < elements[mid] {
			high = mid - 1
		} else {
			if earliestOccurrence == -1 {
				earliestOccurrence = mid
			}
			if mid == len(elements)-1 || elements[mid+1] != elements[mid] {
				return earliestOccurrence, mid
			}
			low = mid + 1
		}
	}
	return -1, -1
}

func firstOccurrence(elements []int, x int, low int, high int) int {
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
