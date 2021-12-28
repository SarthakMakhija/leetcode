package set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes1(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	expected := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	setZeroes(matrix)

	if !reflect.DeepEqual(expected, matrix) {
		t.Fatalf("Expected %v, received %v", expected, matrix)
	}
}

func TestSetZeroes2(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	expected := [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0},
	}
	setZeroes(matrix)

	if !reflect.DeepEqual(expected, matrix) {
		t.Fatalf("Expected %v, received %v", expected, matrix)
	}
}
