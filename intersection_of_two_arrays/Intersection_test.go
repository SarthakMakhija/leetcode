package intersection_of_two_arrays

import "testing"

func TestNumberOfElementsInIntersection1(t *testing.T) {
	arr1 := []int{89, 24, 75, 11, 23}
	arr2 := []int{89, 2, 4}
	result := numberOfElementsInIntersection(arr1, arr2)
	expected := 1

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestNumberOfElementsInIntersection2(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{3, 4, 5, 6, 7}
	result := numberOfElementsInIntersection(arr1, arr2)
	expected := 4

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestNumberOfElementsInIntersection3(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{3, 3, 5, 6, 7}
	result := numberOfElementsInIntersection(arr1, arr2)
	expected := 3

	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
