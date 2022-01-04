package all_duplicates_in_array

import (
	"reflect"
	"testing"
)

func TestFindDuplicates1(t *testing.T) {
	duplicates := findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
	expected := []int{2, 3}

	if !reflect.DeepEqual(expected, duplicates) {
		t.Fatalf("Expected %v, received %v", expected, duplicates)
	}
}

func TestFindDuplicates2(t *testing.T) {
	duplicates := findDuplicates([]int{1, 1, 2})
	expected := []int{1}

	if !reflect.DeepEqual(expected, duplicates) {
		t.Fatalf("Expected %v, received %v", expected, duplicates)
	}
}

func TestFindDuplicates3(t *testing.T) {
	duplicates := findDuplicates([]int{1})
	var expected []int

	if !reflect.DeepEqual(expected, duplicates) {
		t.Fatalf("Expected %v, received %v", expected, duplicates)
	}
}
