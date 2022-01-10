package subarray_with_sum_zero

import (
	"testing"
)

func TestSubArraySumEqualsZero1(t *testing.T) {
	nums := []int{1, 4, 13, -3, -10, 5}
	containsSubArray := checkSubarraySumZero(nums)

	if containsSubArray != true {
		t.Fatalf("Expected subArray with zero sum to be contained but was not")
	}
}

func TestSubArraySumEqualsZero2(t *testing.T) {
	nums := []int{-1, 4, -3, 5, 1}
	containsSubArray := checkSubarraySumZero(nums)

	if containsSubArray != true {
		t.Fatalf("Expected subArray with zero sum to be contained but was not")
	}
}

func TestSubArraySumEqualsZero3(t *testing.T) {
	nums := []int{3, 1, -2, 5, 6}
	containsSubArray := checkSubarraySumZero(nums)

	if containsSubArray != false {
		t.Fatalf("Expected subArray with zero sum to not be contained but was")
	}
}

func TestSubArraySumEqualsZero(t *testing.T) {
	nums := []int{5, 6, 0, 8}
	containsSubArray := checkSubarraySumZero(nums)

	if containsSubArray != true {
		t.Fatalf("Expected subArray with zero sum to be contained but was not")
	}
}
