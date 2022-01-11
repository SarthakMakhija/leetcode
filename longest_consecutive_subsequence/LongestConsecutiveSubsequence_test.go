package longest_consecutive_subsequence

import (
	"testing"
)

func TestLongestConsecutiveSubsequence1(t *testing.T) {
	longest := longestConsecutive([]int{1, 3, 4, 3, 3, 2, 9, 10})
	expected := 4

	if expected != longest {
		t.Fatalf("Expected %v, received %v", expected, longest)
	}
}

func TestLongestConsecutiveSubsequence2(t *testing.T) {
	longest := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	expected := 9

	if expected != longest {
		t.Fatalf("Expected %v, received %v", expected, longest)
	}
}

func TestLongestConsecutiveSubsequence3(t *testing.T) {
	longest := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	expected := 4

	if expected != longest {
		t.Fatalf("Expected %v, received %v", expected, longest)
	}
}
