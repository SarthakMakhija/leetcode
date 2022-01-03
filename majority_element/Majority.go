package majority_element

func majorityElement(nums []int) []int {
	length, frequency, exclude := len(nums), make(map[int]int), make(map[int]bool)
	greaterThanOneThird := 1 + length/3

	var result []int
	for _, num := range nums {
		if exclude[num] {
			continue
		}
		frequency[num] = frequency[num] + 1
		if frequency[num] >= greaterThanOneThird {
			result = append(result, num)
			exclude[num] = true
			delete(frequency, num)
		}
	}
	return result
}
