package combination_sum_4

import (
	"testing"
)

func TestCombinationSum1(t *testing.T) {
	totalCombinations := combinationSum4([]int{2, 3, 6, 7}, 7)
	expected := 4

	if totalCombinations != expected {
		t.Fatalf("Expected %v, received %v", expected, totalCombinations)
	}
}

func TestCombinationSum2(t *testing.T) {
	totalCombinations := combinationSum4([]int{2, 3, 5}, 8)
	expected := 6

	if totalCombinations != expected {
		t.Fatalf("Expected %v, received %v", expected, totalCombinations)
	}
}

func TestCombinationSum3(t *testing.T) {
	totalCombinations := combinationSum4([]int{1, 2, 3}, 4)
	expected := 7

	if totalCombinations != expected {
		t.Fatalf("Expected %v, received %v", expected, totalCombinations)
	}
}
