package count_occurrence_in_sorted

import "testing"

func TestCountOccurrenceInSortedArray1(t *testing.T) {
	elements := []int{1, 3, 5, 5, 5, 5, 67, 123, 125}
	occurrence := countOccurrence(elements, 5)
	expected := 4

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestCountOccurrenceInSortedArray2(t *testing.T) {
	elements := []int{1, 3, 9, 9, 9, 9, 67, 123, 125}
	occurrence := countOccurrence(elements, 5)
	expected := 0

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestCountOccurrenceInSortedArray3(t *testing.T) {
	elements := []int{1, 3, 5, 6, 7, 9, 67, 123, 125}
	occurrence := countOccurrence(elements, 5)
	expected := 1

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestCountOccurrenceInSortedArray4(t *testing.T) {
	elements := []int{1, 3, 5, 6, 7, 9, 9, 123, 125}
	occurrence := countOccurrence(elements, 9)
	expected := 2

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}

func TestCountOccurrenceInSortedArray5(t *testing.T) {
	elements := []int{1, 1, 1, 1, 1, 1}
	occurrence := countOccurrence(elements, 1)
	expected := 6

	if occurrence != expected {
		t.Fatalf("Expected %v, received %v", expected, occurrence)
	}
}
