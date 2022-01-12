package more_than_n_k_occurrence

//moreThanNByK
//One option is to use a hashmap with a count of each element, and anytime the count
//goes beyond n/k, return the element
//Other option is to create a probable set of elements which can be in the output, this solution
//takes O(nk) time
func moreThanNByK(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	nByK := len(nums) / k
	frequency := make(map[int]int)
	for _, num := range nums {
		if _, ok := frequency[num]; ok {
			frequency[num] = frequency[num] + 1
		} else {
			if len(frequency) < k-1 {
				frequency[num] = 1
			} else {
				frequency = reduceByOne(frequency)
			}
		}
	}
	var result []int
	for num := range frequency {
		if count := countOccurrenceOf(num, nums); count >= nByK+1 {
			result = append(result, num)
		}
	}
	return result
}

func countOccurrenceOf(k int, nums []int) int {
	count := 0
	for _, num := range nums {
		if num == k {
			count = count + 1
		}
	}
	return count
}

func reduceByOne(frequency map[int]int) map[int]int {
	for k := range frequency {
		frequency[k] = frequency[k] - 1
		if frequency[k] == 0 {
			delete(frequency, k)
		}
	}
	return frequency
}
