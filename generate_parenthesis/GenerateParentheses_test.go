package generate_parenthesis

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis1(t *testing.T) {
	all := generateParenthesis(3)
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	if !reflect.DeepEqual(expected, all) {
		t.Fatalf("Expected %v, received %v", expected, all)
	}
}

func TestGenerateParenthesis2(t *testing.T) {
	all := generateParenthesis(2)
	expected := []string{"(())", "()()"}

	if !reflect.DeepEqual(expected, all) {
		t.Fatalf("Expected %v, received %v", expected, all)
	}
}

func TestGenerateParenthesis3(t *testing.T) {
	all := generateParenthesis(1)
	expected := []string{"()"}

	if !reflect.DeepEqual(expected, all) {
		t.Fatalf("Expected %v, received %v", expected, all)
	}
}
