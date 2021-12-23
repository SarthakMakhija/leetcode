package minimum_path_sum

import "testing"

func TestMinimumPathSum1(t *testing.T) {
	sum := minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	})
	expected := 7

	if expected != sum {
		t.Fatalf("Expected sum %v, received %v", expected, sum)
	}
}

func TestMinimumPathSum2(t *testing.T) {
	sum := minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	})
	expected := 12

	if expected != sum {
		t.Fatalf("Expected sum %v, received %v", expected, sum)
	}
}

func TestMinimumPathSum3(t *testing.T) {
	sum := minPathSum([][]int{
		{3, 2, 12, 15, 10},
		{6, 19, 7, 11, 17},
		{8, 5, 12, 32, 21},
		{3, 20, 2, 9, 7},
	})
	expected := 52

	if expected != sum {
		t.Fatalf("Expected sum %v, received %v", expected, sum)
	}
}
