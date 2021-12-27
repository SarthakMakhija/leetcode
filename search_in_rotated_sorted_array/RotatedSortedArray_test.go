package search_in_rotated_sorted_array

import (
	"testing"
)

func TestSearchInUnRotatedArray1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	index := search(nums, 5)

	if index != 4 {
		t.Fatalf("Expected index to be 4 but was %v", index)
	}
}

func TestSearchInUnRotatedArray2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	index := search(nums, 2)

	if index != 1 {
		t.Fatalf("Expected index to be 1 but was %v", index)
	}
}

func TestSearchInUnRotatedArray3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	index := search(nums, 10)

	if index != -1 {
		t.Fatalf("Expected index to be -1 but was %v", index)
	}
}

func TestSearchInRotatedArray1(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 14, 0, 1, 2}
	index := search(nums, 10)

	if index != 6 {
		t.Fatalf("Expected index to be 6 but was %v", index)
	}
}

func TestSearchInRotatedArray2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 14, 0, 1, 2}
	index := search(nums, 8)

	if index != 4 {
		t.Fatalf("Expected index to be 4 but was %v", index)
	}
}

func TestSearchInRotatedArray3(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 14, 0, 1, 2}
	index := search(nums, 14)

	if index != 8 {
		t.Fatalf("Expected index to be 8 but was %v", index)
	}
}

func TestSearchInRotatedArray4(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 14, 0, 1, 2}
	index := search(nums, 0)

	if index != 9 {
		t.Fatalf("Expected index to be 9 but was %v", index)
	}
}

func TestSearchInRotatedArray5(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 14, 0, 1, 2}
	index := search(nums, 2)

	if index != 11 {
		t.Fatalf("Expected index to be 11 but was %v", index)
	}
}

func TestSearchInRotatedArray6(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 14, 0, 1, 2}
	index := search(nums, 20)

	if index != -1 {
		t.Fatalf("Expected index to be -1 but was %v", index)
	}
}

func TestSearchInRotatedArray7(t *testing.T) {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	index := search(nums, 9)

	if index != 7 {
		t.Fatalf("Expected index to be 7 but was %v", index)
	}
}
