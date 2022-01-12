package union_of_two_arrays

import "testing"

func TestUnion1(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3}

	unionCount := doUnion(arr1, arr2)
	expected := 5

	if unionCount != expected {
		t.Fatalf("Expected %v, received %v", expected, unionCount)
	}
}

func TestUnion2(t *testing.T) {
	arr1 := []int{85, 25, 1, 32, 54, 6}
	arr2 := []int{85, 2}

	unionCount := doUnion(arr1, arr2)
	expected := 7

	if unionCount != expected {
		t.Fatalf("Expected %v, received %v", expected, unionCount)
	}
}

func TestUnion3(t *testing.T) {
	arr1 := []int{85, 25, 1, 32, 54, 6}
	arr2 := []int{85, 25, 1, 32, 54, 6}

	unionCount := doUnion(arr1, arr2)
	expected := 6

	if unionCount != expected {
		t.Fatalf("Expected %v, received %v", expected, unionCount)
	}
}
