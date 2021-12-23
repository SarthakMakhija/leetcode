package unique_paths

import "testing"

func TestUniquePaths1(t *testing.T) {
	pathCount := uniquePaths(3, 2)
	if pathCount != 3 {
		t.Fatalf("Expected path count to be %v, received %v", 3, pathCount)
	}
}

func TestUniquePaths2(t *testing.T) {
	pathCount := uniquePaths(3, 7)
	if pathCount != 28 {
		t.Fatalf("Expected path count to be %v, received %v", 28, pathCount)
	}
}
