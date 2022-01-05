package unique_paths_2

import "testing"

func TestUniquePathsWithObstacles1(t *testing.T) {
	pathCount := uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	if pathCount != 2 {
		t.Fatalf("Expected %v, received %v", 2, pathCount)
	}
}

func TestUniquePathsWithObstacles2(t *testing.T) {
	pathCount := uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0},
	})
	if pathCount != 1 {
		t.Fatalf("Expected %v, received %v", 1, pathCount)
	}
}

func TestUniquePathsWithObstacles3(t *testing.T) {
	pathCount := uniquePathsWithObstacles([][]int{
		{0, 1},
		{1, 0},
	})
	if pathCount != 0 {
		t.Fatalf("Expected %v, received %v", 0, pathCount)
	}
}

func TestUniquePathsWithObstacles4(t *testing.T) {
	pathCount := uniquePathsWithObstacles([][]int{
		{0, 0},
		{0, 1},
	})
	if pathCount != 0 {
		t.Fatalf("Expected %v, received %v", 0, pathCount)
	}
}
