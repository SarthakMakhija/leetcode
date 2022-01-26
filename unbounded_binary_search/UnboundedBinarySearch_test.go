package unbounded_binary_search

import "testing"

func TestUnboundedBinarySearch1(t *testing.T) {
	elements := []int{1, 10, 15, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200}
	index := search(elements, 100)
	expected := 7

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestUnboundedBinarySearch2(t *testing.T) {
	elements := []int{1, 10, 15, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200}
	index := search(elements, 200)
	expected := 12

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestUnboundedBinarySearch3(t *testing.T) {
	elements := []int{1, 10, 15, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200}
	index := search(elements, 1)
	expected := 0

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestUnboundedBinarySearch4(t *testing.T) {
	elements := []int{1, 10, 15, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200}
	index := search(elements, 900)
	expected := -1

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}

func TestUnboundedBinarySearch5(t *testing.T) {
	elements := []int{1, 10, 15, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200}
	index := search(elements, 0)
	expected := -1

	if index != expected {
		t.Fatalf("Expected %v, received %v", expected, index)
	}
}
