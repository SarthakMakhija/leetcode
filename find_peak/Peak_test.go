package find_peak

import (
	"testing"
)

func TestFindPeak1(t *testing.T) {
	index := findPeakElement([]int{1, 2, 3, 1})
	expected := 2

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestFindPeak2(t *testing.T) {
	index := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
	expected := 5

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestFindPeak3(t *testing.T) {
	index := findPeakElement([]int{1, 2})
	expected := 1

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestFindPeak4(t *testing.T) {
	index := findPeakElement([]int{2, 1})
	expected := 0

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}
