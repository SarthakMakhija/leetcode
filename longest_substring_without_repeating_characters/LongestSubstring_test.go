package longest_substring_without_repeating_characters

import (
	"testing"
)

func TestLongestSubstringLength1(t *testing.T) {
	length := lengthOfLongestSubstring("abcabcbb")
	expected := 3

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}

func TestLongestSubstringLength2(t *testing.T) {
	length := lengthOfLongestSubstring("bbbb")
	expected := 1

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}

func TestLongestSubstringLength3(t *testing.T) {
	length := lengthOfLongestSubstring("pwwkew")
	expected := 3

	if length != expected {
		t.Fatalf("Expected %v, received %v", expected, length)
	}
}
