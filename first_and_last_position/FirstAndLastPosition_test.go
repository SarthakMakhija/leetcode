package first_and_last_position

import (
	"reflect"
	"testing"
)

func TestFirstAndLastPosition1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 8)
	expected := []int{3, 4}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 7)
	expected := []int{1, 2}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition3(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 5)
	expected := []int{0, 0}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition4(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 10)
	expected := []int{5, 5}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition5(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 20)
	expected := []int{-1, -1}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition6(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 6)
	expected := []int{-1, -1}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition7(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	indices := searchRange(nums, 9)
	expected := []int{-1, -1}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition8(t *testing.T) {
	nums := []int{1, 3}
	indices := searchRange(nums, 1)
	expected := []int{0, 0}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition9(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 4, 5, 5}
	indices := searchRange(nums, 2)
	expected := []int{-1, -1}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}

func TestFirstAndLastPosition10(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 4, 5, 5}
	indices := searchRange(nums, 2)
	expected := []int{-1, -1}

	if !reflect.DeepEqual(expected, indices) {
		t.Fatalf("Expected %v, received %v", expected, indices)
	}
}
