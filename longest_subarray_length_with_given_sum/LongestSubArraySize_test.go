package longest_subarray_length_with_given_sum

import (
	"testing"
)

func TestLongestSubArrayLengthWithSumK1(t *testing.T) {
	nums := []int{5, 8, -4, -4, 9, -2, 2}
	length := longestSubArrayLengthWithSum(nums, 0)

	if length != 3 {
		t.Fatalf("Expected longest subarray length to be %v, received %v", 3, length)
	}
}

func TestLongestSubArrayLengthWithSumK2(t *testing.T) {
	nums := []int{3, 1, 0, 1, 8, 2, 3, 6}
	length := longestSubArrayLengthWithSum(nums, 5)

	if length != 4 {
		t.Fatalf("Expected longest subarray length to be %v, received %v", 4, length)
	}
}

func TestLongestSubArrayLengthWithSumK3(t *testing.T) {
	nums := []int{8, 3, 7}
	length := longestSubArrayLengthWithSum(nums, 15)

	if length != 0 {
		t.Fatalf("Expected longest subarray length to be %v, received %v", 0, length)
	}
}

func TestLongestSubArrayLengthWithSumK4(t *testing.T) {
	nums := []int{10, 5, 2, 7, 1, 9}
	length := longestSubArrayLengthWithSum(nums, 15)

	if length != 4 {
		t.Fatalf("Expected longest subarray length to be %v, received %v", 4, length)
	}
}

func TestLongestSubArrayLengthWithSumK5(t *testing.T) {
	nums := []int{1, 4, 3, 3, 5, 5}
	length := longestSubArrayLengthWithSum(nums, 15)

	if length != 4 {
		t.Fatalf("Expected longest subarray length to be %v, received %v", 4, length)
	}
}

func TestLongestSubArrayLengthWithSumK6(t *testing.T) {
	nums := []int{-13, 0, 6, 15, 16, 2, 15, -12, 17, -16, 0, -3, 19, -3, 2, -9, -6}
	length := longestSubArrayLengthWithSum(nums, 15)

	if length != 5 {
		t.Fatalf("Expected longest subarray length to be %v, received %v", 5, length)
	}
}
