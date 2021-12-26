package triangle

import (
	"testing"
)

func TestMinimumTriangleSum1(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	sum := minimumTotal(triangle)
	expected := 11

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestMinimumTriangleSum2(t *testing.T) {
	triangle := [][]int{
		{2},
		{1, -6},
		{6, 1, 7},
		{4, 7, 8, 3},
	}
	sum := minimumTotal(triangle)
	expected := 4

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestMinimumTriangleSum3(t *testing.T) {
	triangle := [][]int{
		{-10},
	}
	sum := minimumTotal(triangle)
	expected := -10

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}
