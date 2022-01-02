package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors1(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	expected := []int{0, 0, 1, 1, 2, 2}

	if !reflect.DeepEqual(expected, nums) {
		t.Fatalf("Expected %v, receivede %v", expected, nums)
	}
}

func TestSortColors2(t *testing.T) {
	nums := []int{2, 0, 1}
	sortColors(nums)
	expected := []int{0, 1, 2}

	if !reflect.DeepEqual(expected, nums) {
		t.Fatalf("Expected %v, receivede %v", expected, nums)
	}
}

func TestSortColors3(t *testing.T) {
	nums := []int{1, 1, 2, 2, 0, 0}
	sortColors(nums)
	expected := []int{0, 0, 1, 1, 2, 2}

	if !reflect.DeepEqual(expected, nums) {
		t.Fatalf("Expected %v, receivede %v", expected, nums)
	}
}
