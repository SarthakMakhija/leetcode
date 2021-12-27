package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	result := twoSum([]int{2, 7, 1, 15}, 9)
	expected := []int{0, 1}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestTwoSum2(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestTwoSum3(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	expected := []int{0, 1}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestTwoSum4(t *testing.T) {
	result := twoSum([]int{3, 4, 5}, 10)
	var expected []int

	if len(result) != 0 {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
