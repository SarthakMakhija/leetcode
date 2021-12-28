package zigzag_conversion

import (
	"testing"
)

func TestConvert1(t *testing.T) {
	str := "PAYPALISHIRING"
	result := convert(str, 3)
	expected := "PAHNAPLSIIGYIR"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestConvert2(t *testing.T) {
	str := "PAYPALISHIRING"
	result := convert(str, 4)
	expected := "PINALSIGYAHRPI"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestConvert3(t *testing.T) {
	str := "PAYPALISHIRING"
	result := convert(str, 1)
	expected := "PAYPALISHIRING"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestConvert4(t *testing.T) {
	str := "AB"
	result := convert(str, 3)
	expected := "AB"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestConvert5(t *testing.T) {
	str := "PAYPALISHIRING"
	result := convert(str, 5)
	expected := "PHASIYIRPLIGAN"

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
