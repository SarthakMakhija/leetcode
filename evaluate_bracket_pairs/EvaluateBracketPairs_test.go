package evaluate_bracket_pairs

import (
	"testing"
)

func TestEvaluateBracketPairs1(t *testing.T) {
	result := evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "12"}})
	expected := "bobis12yearsold"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestEvaluateBracketPairs2(t *testing.T) {
	result := evaluate("hi(name)", [][]string{{"age", "12"}})
	expected := "hi?"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestEvaluateBracketPairs3(t *testing.T) {
	result := evaluate("(a)(a)(a)aaa", [][]string{{"a", "yes"}, {"age", "12"}})
	expected := "yesyesyesaaa"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
