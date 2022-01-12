package sort_an_array_according_to_other

import (
	"reflect"
	"testing"
)

func TestRelativeSorted1(t *testing.T) {
	A1 := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8}
	A2 := []int{2, 1, 8, 3}

	sorted := sortA1ByA2(A1, A2)
	expected := []int{2, 2, 1, 1, 8, 8, 3, 5, 6, 7, 9}

	if !reflect.DeepEqual(expected, sorted) {
		t.Fatalf("Expected %v, received %v", expected, sorted)
	}
}

func TestRelativeSorted2(t *testing.T) {
	A1 := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8, 7}
	A2 := []int{2, 1, 8, 3}

	sorted := sortA1ByA2(A1, A2)
	expected := []int{2, 2, 1, 1, 8, 8, 3, 5, 6, 7, 7, 9}

	if !reflect.DeepEqual(expected, sorted) {
		t.Fatalf("Expected %v, received %v", expected, sorted)
	}
}

func TestRelativeSorted3(t *testing.T) {
	A1 := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8}
	A2 := []int{99, 22, 444, 56}

	sorted := sortA1ByA2(A1, A2)
	expected := []int{1, 1, 2, 2, 3, 5, 6, 7, 8, 8, 9}

	if !reflect.DeepEqual(expected, sorted) {
		t.Fatalf("Expected %v, received %v", expected, sorted)
	}
}
