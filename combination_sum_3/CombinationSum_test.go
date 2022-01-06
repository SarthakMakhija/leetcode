package combination_sum_3

import (
	"reflect"
	"testing"
)

func TestCombinationSum3_1(t *testing.T) {
	combinations := combinationSum3(3, 7)
	expected := [][]int{{1, 2, 4}}

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}

func TestCombinationSum3_2(t *testing.T) {
	combinations := combinationSum3(3, 9)
	expected := [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}

func TestCombinationSum3_3(t *testing.T) {
	combinations := combinationSum3(4, 1)
	var expected [][]int

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}

func TestCombinationSum3_4(t *testing.T) {
	combinations := combinationSum3(2, 18)
	var expected [][]int

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}
