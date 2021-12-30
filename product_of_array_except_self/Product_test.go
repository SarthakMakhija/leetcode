package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf1(t *testing.T) {
	product := productExceptSelf([]int{1, 2, 3, 4})
	expected := []int{24, 12, 8, 6}

	if !reflect.DeepEqual(expected, product) {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}

func TestProductExceptSelf2(t *testing.T) {
	product := productExceptSelf([]int{-1, 1, 0, -3, 3})
	expected := []int{0, 0, 9, 0, 0}

	if !reflect.DeepEqual(expected, product) {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}

func TestProductExceptSelf3(t *testing.T) {
	product := productExceptSelf([]int{6, 2, 3, 4, 5})
	expected := []int{120, 360, 240, 180, 144}

	if !reflect.DeepEqual(expected, product) {
		t.Fatalf("Expected %v, received %v", expected, product)
	}
}
