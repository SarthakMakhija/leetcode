package unbounded_binary_search

func search(elements []int, x int) int {
	if len(elements) == 0 {
		return -1
	}
	if elements[0] == x {
		return 0
	}
	index, length := 1, len(elements)
	if x > elements[length-1] {
		return -1
	}
	for index < length-1 && elements[index] < x {
		index = 2 * index
	}
	if index < length-1 {
		if elements[index] == x {
			return index
		}
		return binarySearch(elements, x, index/2+1, index)
	}
	return binarySearch(elements, x, index/2+1, index)
}

func binarySearch(elements []int, x int, low int, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if x < elements[mid] {
			high = mid - 1
		} else if x > elements[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
