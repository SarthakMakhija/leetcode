package remove_duplicates_from_sorted_array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	length := removeDuplicates(nums)
	expectedElements, expectedLength := []int{1, 2, 3, 4, 5}, 5

	if length != expectedLength {
		t.Fatalf("Expected length to be %v, received %v", expectedLength, length)
	}

	if !reflect.DeepEqual(expectedElements, nums[0:length]) {
		t.Fatalf("Expected elements to be %v, received %v", expectedElements, nums[0:length])
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}
	length := removeDuplicates(nums)
	expectedElements, expectedLength := []int{1, 1, 2, 2, 3, 3}, 6

	if length != expectedLength {
		t.Fatalf("Expected length to be %v, received %v", expectedLength, length)
	}

	if !reflect.DeepEqual(expectedElements, nums[0:length]) {
		t.Fatalf("Expected elements to be %v, received %v", expectedElements, nums[0:length])
	}
}
func TestRemoveDuplicates3(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	length := removeDuplicates(nums)
	expectedElements, expectedLength := []int{0, 0, 1, 1, 2, 3, 3}, 7

	if length != expectedLength {
		t.Fatalf("Expected length to be %v, received %v", expectedLength, length)
	}

	if !reflect.DeepEqual(expectedElements, nums[0:length]) {
		t.Fatalf("Expected elements to be %v, received %v", expectedElements, nums[0:length])
	}
}

func TestRemoveDuplicates4(t *testing.T) {
	nums := []int{0, 1, 3, 3, 3, 3}
	length := removeDuplicates(nums)
	expectedElements, expectedLength := []int{0, 1, 3, 3}, 4

	if length != expectedLength {
		t.Fatalf("Expected length to be %v, received %v", expectedLength, length)
	}

	if !reflect.DeepEqual(expectedElements, nums[0:length]) {
		t.Fatalf("Expected elements to be %v, received %v", expectedElements, nums[0:length])
	}
}
