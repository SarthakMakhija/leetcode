package first_repeating_element

import "testing"

func TestFirstRepeatingElement1(t *testing.T) {
	index := firstRepeated([]int{1, 2, 3, 4})
	expected := -1

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestFirstRepeatingElement2(t *testing.T) {
	index := firstRepeated([]int{1, 5, 3, 4, 3, 5, 6})
	expected := 2

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}
