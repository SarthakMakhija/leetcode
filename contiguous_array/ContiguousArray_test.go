package contiguous_array

import (
	"testing"
)

func TestLongestSubArrayWithEqualZerosAndOnes1(t *testing.T) {
	nums := []int{1, 0, 1, 1, 1, 0, 0}
	length := findMaxLength(nums)

	if length != 6 {
		t.Fatalf("Expected LongestSubArrayWithEqualZerosAndOnes to be 6, received %v", length)
	}
}

func TestLongestSubArrayWithEqualZerosAndOnes2(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	length := findMaxLength(nums)

	if length != 0 {
		t.Fatalf("Expected LongestSubArrayWithEqualZerosAndOnes to be 0, received %v", length)
	}
}

func TestLongestSubArrayWithEqualZerosAndOnes3(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1}
	length := findMaxLength(nums)

	if length != 4 {
		t.Fatalf("Expected LongestSubArrayWithEqualZerosAndOnes to be 4, received %v", length)
	}
}

func TestLongestSubArrayWithEqualZerosAndOnes4(t *testing.T) {
	nums := []int{0, 1, 0}
	length := findMaxLength(nums)

	if length != 2 {
		t.Fatalf("Expected LongestSubArrayWithEqualZerosAndOnes to be 2, received %v", length)
	}
}
