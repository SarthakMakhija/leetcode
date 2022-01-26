package sqrt

import "testing"

func TestSqrt1(t *testing.T) {
	result := sqrt(25)
	expected := 5

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSqrt2(t *testing.T) {
	result := sqrt(36)
	expected := 6

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSqrt3(t *testing.T) {
	result := sqrt(26)
	expected := 5

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSqrt4(t *testing.T) {
	result := sqrt(49)
	expected := 7

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSqrt5(t *testing.T) {
	result := sqrt(48)
	expected := 6

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSqrt6(t *testing.T) {
	result := sqrt(1)
	expected := 1

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSqrt7(t *testing.T) {
	result := sqrt(0)
	expected := 0

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
