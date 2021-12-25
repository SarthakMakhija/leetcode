package rpn

import "testing"

func TestEvaluateRPN1(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	result := evalRPN(tokens)
	expected := 9

	if result != expected {
		t.Fatalf("Expectede %v, received %v", expected, result)
	}
}

func TestEvaluateRPN2(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+"}
	result := evalRPN(tokens)
	expected := 6

	if result != expected {
		t.Fatalf("Expectede %v, received %v", expected, result)
	}
}

func TestEvaluateRPN3(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	result := evalRPN(tokens)
	expected := 22

	if result != expected {
		t.Fatalf("Expectede %v, received %v", expected, result)
	}
}
