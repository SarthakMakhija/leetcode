package sqrt

func sqrt(x int) int {
	if x < 0 {
		panic("Negative number")
	}
	low, high := 1, x
	result := 0
	for low <= high {
		mid := (low + high) / 2
		square := mid * mid
		if square == x {
			return mid
		}
		if x < square {
			high = mid - 1
		} else {
			low = mid + 1
			result = mid
		}
	}
	return result
}
