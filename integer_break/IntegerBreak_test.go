package integer_break

import (
	"testing"
)

func TestIntegerBreak1(t *testing.T) {
	product := integerBreak(3)
	expected := 2

	if expected != product {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}

func TestIntegerBreak2(t *testing.T) {
	product := integerBreak(2)
	expected := 1

	if expected != product {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}

func TestIntegerBreak3(t *testing.T) {
	product := integerBreak(10)
	expected := 36

	if expected != product {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}

func TestIntegerBreak4(t *testing.T) {
	product := integerBreak(5)
	expected := 6

	if expected != product {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}
