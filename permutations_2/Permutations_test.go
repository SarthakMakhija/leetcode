package permutations_2

import (
	"reflect"
	"testing"
)

func TestPermutations1(t *testing.T) {
	permutations := permuteUnique([]int{1, 2, 3})
	expected := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}

	if !reflect.DeepEqual(expected, permutations) {
		t.Fatalf("Expected %v, received %v", expected, permutations)
	}
}

func TestPermutations2(t *testing.T) {
	permutations := permuteUnique([]int{1, 1, 2})
	expected := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}

	if !reflect.DeepEqual(expected, permutations) {
		t.Fatalf("Expected %v, received %v", expected, permutations)
	}
}

func TestPermutations3(t *testing.T) {
	permutations := permuteUnique([]int{1, 1})
	expected := [][]int{{1, 1}}

	if !reflect.DeepEqual(expected, permutations) {
		t.Fatalf("Expected %v, received %v", expected, permutations)
	}
}
