package maximum_subarray_sum

import (
	"testing"
)

func TestMaxSubArraySum1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArraySum := maxSubArray(nums)
	expected := 6

	if expected != maxSubArraySum {
		t.Fatalf("Expected %v, received %v", expected, maxSubArraySum)
	}
}

func TestMaxSubArraySum2(t *testing.T) {
	nums := []int{-2, 3, 2, -1}
	maxSubArraySum := maxSubArray(nums)
	expected := 5

	if expected != maxSubArraySum {
		t.Fatalf("Expected %v, received %v", expected, maxSubArraySum)
	}
}

func TestMaxSubArraySum3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	maxSubArraySum := maxSubArray(nums)
	expected := 23

	if expected != maxSubArraySum {
		t.Fatalf("Expected %v, received %v", expected, maxSubArraySum)
	}
}
