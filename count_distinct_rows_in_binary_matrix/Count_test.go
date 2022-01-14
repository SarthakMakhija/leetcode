package count_distinct_rows_in_binary_matrix

import (
	"testing"
)

func TestCountDistinctRowsInBinaryMatrix1(t *testing.T) {
	matrix := [][]bool{
		{true, false, false},
		{true, true, true},
		{true, false, false},
		{true, true, true},
	}
	count := countDistinctRowsIn(matrix)
	expected := 2

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestCountDistinctRowsInBinaryMatrix2(t *testing.T) {
	matrix := [][]bool{
		{true, false, false},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	}
	count := countDistinctRowsIn(matrix)
	expected := 3

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestCountDistinctRowsInBinaryMatrix3(t *testing.T) {
	matrix := [][]bool{
		{true, true, false, false},
		{true, true, false, false},
		{true, true, false, false},
		{true, true, false, false},
	}
	count := countDistinctRowsIn(matrix)
	expected := 1

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestCountDistinctRowsInBinaryMatrix4(t *testing.T) {
	matrix := [][]bool{
		{true, true, false, false},
		{true, false, true, true},
		{true, true, false, true},
		{false, false, false, false},
	}
	count := countDistinctRowsIn(matrix)
	expected := 4

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}
