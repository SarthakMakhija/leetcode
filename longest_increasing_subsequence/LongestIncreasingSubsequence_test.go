package longest_increasing_subsequence

import (
	"testing"
)

func TestLongestIncreasingSubsequenceLength1(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected := 4
	length := lengthOfLIS(nums)

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}

func TestLongestIncreasingSubsequenceLength2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	expected := 4
	length := lengthOfLIS(nums)

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}

func TestLongestIncreasingSubsequenceLength3(t *testing.T) {
	nums := []int{7, 5, 2, 4, 7, 2, 3, 6, 4, 5, 12, 1, 7}
	expected := 5
	length := lengthOfLIS(nums)

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}

func TestLongestIncreasingSubsequenceLength4(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	expected := 1
	length := lengthOfLIS(nums)

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}
