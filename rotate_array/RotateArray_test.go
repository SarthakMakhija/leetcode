package rotate_array

import (
	"reflect"
	"testing"
)

func TestRotateRightBy3_1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	expected := []int{5, 6, 7, 1, 2, 3, 4}

	if !reflect.DeepEqual(expected, nums) {
		t.Fatalf("Expected %v, received %v", expected, nums)
	}
}

func TestRotateRightBy3_2(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
	expected := []int{3, 99, -1, -100}

	if !reflect.DeepEqual(expected, nums) {
		t.Fatalf("Expected %v, received %v", expected, nums)
	}
}
