package find_minimum_in_rotated_sorted_array

import (
	"testing"
)

func TestFindMinimum1(t *testing.T) {
	minimum := findMin([]int{3, 4, 5, 1, 2})
	expected := 1

	if expected != minimum {
		t.Fatalf("Expected %v, received %v", expected, minimum)
	}
}

func TestFindMinimum2(t *testing.T) {
	minimum := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	expected := 0

	if expected != minimum {
		t.Fatalf("Expected %v, received %v", expected, minimum)
	}
}

func TestFindMinimum3(t *testing.T) {
	minimum := findMin([]int{11, 13, 15, 17})
	expected := 11

	if expected != minimum {
		t.Fatalf("Expected %v, received %v", expected, minimum)
	}
}

func TestFindMinimum4(t *testing.T) {
	minimum := findMin([]int{4, 5, 6, 7, 8, 9, 10, -2, -1, 0, 1, 2, 3})
	expected := -2

	if expected != minimum {
		t.Fatalf("Expected %v, received %v", expected, minimum)
	}
}

func TestFindMinimum5(t *testing.T) {
	minimum := findMin([]int{2, 1})
	expected := 1

	if expected != minimum {
		t.Fatalf("Expected %v, received %v", expected, minimum)
	}
}

func TestFindMinimum6(t *testing.T) {
	minimum := findMin([]int{2})
	expected := 2

	if expected != minimum {
		t.Fatalf("Expected %v, received %v", expected, minimum)
	}
}
