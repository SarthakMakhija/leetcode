package word_break

import (
	"testing"
)

func TestWordBreak1(t *testing.T) {
	possible := wordBreak("leetcode", []string{"leet", "code"})
	if possible != true {
		t.Fatalf("Expected leetcode to be broken based on dictionary")
	}
}

func TestWordBreak2(t *testing.T) {
	possible := wordBreak("applepenapple", []string{"apple", "pen"})
	if possible != true {
		t.Fatalf("Expected applepenapple to be broken based on dictionary")
	}
}

func TestWordBreak3(t *testing.T) {
	possible := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	if possible != false {
		t.Fatalf("Expected catsandog to not be broken based on dictionary")
	}
}

func TestWordBreak4(t *testing.T) {
	possible := wordBreak("abcd", []string{"a", "abc", "b", "cd"})
	if possible != true {
		t.Fatalf("Expected abcd to be broken based on dictionary")
	}
}
