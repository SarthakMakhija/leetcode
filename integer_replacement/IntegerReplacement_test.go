package integer_replacement

import (
	"testing"
)

func TestIntegerReplacement1(t *testing.T) {
	steps := integerReplacement(4)
	expected := 2

	if steps != expected {
		t.Fatalf("Expected %v, received %v", expected, steps)
	}
}

func TestIntegerReplacement2(t *testing.T) {
	steps := integerReplacement(8)
	expected := 3

	if steps != expected {
		t.Fatalf("Expected %v, received %v", expected, steps)
	}
}

func TestIntegerReplacement3(t *testing.T) {
	steps := integerReplacement(7)
	expected := 4

	if steps != expected {
		t.Fatalf("Expected %v, received %v", expected, steps)
	}
}

func TestIntegerReplacement4(t *testing.T) {
	steps := integerReplacement(11)
	expected := 5

	if steps != expected {
		t.Fatalf("Expected %v, received %v", expected, steps)
	}
}

func TestIntegerReplacement5(t *testing.T) {
	steps := integerReplacement(21)
	expected := 6

	if steps != expected {
		t.Fatalf("Expected %v, received %v", expected, steps)
	}
}
