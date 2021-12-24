package decode_ways

import (
	"testing"
)

func TestNumDecodings1(t *testing.T) {
	str := "226"
	count := numDecodings(str)
	expected := 3

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings2(t *testing.T) {
	str := "224"
	count := numDecodings(str)
	expected := 3

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings3(t *testing.T) {
	str := "12"
	count := numDecodings(str)
	expected := 2

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings4(t *testing.T) {
	str := "10"
	count := numDecodings(str)
	expected := 1

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings5(t *testing.T) {
	str := "06"
	count := numDecodings(str)
	expected := 0

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings6(t *testing.T) {
	str := "11106"
	count := numDecodings(str)
	expected := 2

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings7(t *testing.T) {
	str := "27"
	count := numDecodings(str)
	expected := 1

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}

func TestNumDecodings8(t *testing.T) {
	str := "123123"
	count := numDecodings(str)
	expected := 9

	if count != expected {
		t.Fatalf("Expected %v, received %v", expected, count)
	}
}
