package longest_common_span_with_same_sum

import (
	"testing"
)

func TestLongestCommonSum1(t *testing.T) {
	longestCommonSum := longestCommonSum([]bool{false, true, false, false, false, false}, []bool{true, false, true, false, false, true})
	expected := 4

	if longestCommonSum != expected {
		t.Fatalf("Expected %v, received %v", expected, longestCommonSum)
	}
}

func TestLongestCommonSum2(t *testing.T) {
	longestCommonSum := longestCommonSum([]bool{false, true, false}, []bool{true, false, true})
	expected := 2

	if longestCommonSum != expected {
		t.Fatalf("Expected %v, received %v", expected, longestCommonSum)
	}
}
