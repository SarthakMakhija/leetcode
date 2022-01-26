package index_of_first_occurrence_in_sorted

import "testing"

func TestFirstOccurrenceInSortedArray1(t *testing.T) {
	elements := []int{1, 3, 5, 5, 5, 5, 67, 123, 125}
	occurrence := firstOccurrence(elements, 5)
	expected := 2

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestFirstOccurrenceInSortedArray2(t *testing.T) {
	elements := []int{1, 3, 9, 9, 9, 9, 67, 123, 125}
	occurrence := firstOccurrence(elements, 5)
	expected := -1

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestFirstOccurrenceInSortedArray3(t *testing.T) {
	elements := []int{1, 3, 5, 6, 7, 9, 67, 123, 125}
	occurrence := firstOccurrence(elements, 5)
	expected := 2

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestFirstOccurrenceInSortedArray4(t *testing.T) {
	elements := []int{1, 3, 5, 6, 7, 9, 9, 123, 125}
	occurrence := firstOccurrence(elements, 9)
	expected := 5

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestFirstOccurrenceInSortedArray5(t *testing.T) {
	elements := []int{1, 1, 1, 1, 1, 1}
	occurrence := firstOccurrence(elements, 1)
	expected := 0

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}
