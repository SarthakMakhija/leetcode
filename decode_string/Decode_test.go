package decode_string

import (
	"testing"
)

func TestDecode1(t *testing.T) {
	str := "3[a]2[bc]"
	decoded := decodeString(str)
	expected := "aaabcbc"

	if decoded != expected {
		t.Fatalf("Expected %v, received %v", expected, decoded)
	}
}

func TestDecode2(t *testing.T) {
	str := "2[abc]3[cd]ef"
	decoded := decodeString(str)
	expected := "abcabccdcdcdef"

	if decoded != expected {
		t.Fatalf("Expected %v, received %v", expected, decoded)
	}
}

func TestDecode3(t *testing.T) {
	str := "3[a2[c]]"
	decoded := decodeString(str)
	expected := "accaccacc"

	if decoded != expected {
		t.Fatalf("Expected %v, received %v", expected, decoded)
	}
}

func TestDecode4(t *testing.T) {
	str := "5[leetcode]"
	decoded := decodeString(str)
	expected := "leetcodeleetcodeleetcodeleetcodeleetcode"

	if decoded != expected {
		t.Fatalf("Expected %v, received %v", expected, decoded)
	}
}
