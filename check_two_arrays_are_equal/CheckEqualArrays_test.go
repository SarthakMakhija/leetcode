package check_two_arrays_are_equal

import "testing"

func TestTwoArraysEqual1(t *testing.T) {
	arr1 := []int{1, 2, 5, 4, 0}
	arr2 := []int{2, 4, 5, 0, 1}

	equal := check(arr1, arr2)
	if equal != true {
		t.Fatalf("Expected arrays to be equal but were not")
	}
}

func TestTwoArraysEqual2(t *testing.T) {
	arr1 := []int{1, 2, 5}
	arr2 := []int{2, 4, 15}

	equal := check(arr1, arr2)
	if equal != false {
		t.Fatalf("Expected arrays to be unequal but were equal")
	}
}

func TestTwoArraysEqual3(t *testing.T) {
	arr1 := []int{1, 2, 2}
	arr2 := []int{2, 2, 1}

	equal := check(arr1, arr2)
	if equal != true {
		t.Fatalf("Expected arrays to be equal but were not")
	}
}

func TestTwoArraysEqual4(t *testing.T) {
	arr1 := []int{3, 5, 6, 6, 2, 8, 2, 12, 16}
	arr2 := []int{12, 8, 5, 16, 10, 19, 6, 2, 3}

	equal := check(arr1, arr2)
	if equal != false {
		t.Fatalf("Expected arrays to be unequal but were equal")
	}
}
