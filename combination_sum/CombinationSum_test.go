package combination_sum

import (
	"reflect"
	"testing"
)

func TestCombinationSum1(t *testing.T) {
	uniqueCombinations := combinationSum([]int{2, 3, 6, 7}, 7)
	expected := [][]int{{3, 2, 2}, {7}}

	if !reflect.DeepEqual(expected, uniqueCombinations) {
		t.Fatalf("Expected %v, received %v", expected, uniqueCombinations)
	}
}

func TestCombinationSum2(t *testing.T) {
	uniqueCombinations := combinationSum([]int{2, 3, 5}, 8)
	expected := [][]int{{2, 2, 2, 2}, {3, 3, 2}, {5, 3}}

	if !reflect.DeepEqual(expected, uniqueCombinations) {
		t.Fatalf("Expected %v, received %v", expected, uniqueCombinations)
	}
}
