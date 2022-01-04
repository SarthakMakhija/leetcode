package all_duplicates_in_array

func findDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var result []int
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num] = frequency[num] + 1
		if frequency[num] == 2 {
			result = append(result, num)
		}
	}
	return result
}
