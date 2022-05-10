package product_of_last_k

import (
	"reflect"
	"testing"
)

func TestProductOfNumbers1(t *testing.T) {
	productOfNumbers := Constructor()
	productOfNumbers.Add(2)
	productOfNumbers.Add(2)
	productOfNumbers.Add(3)
	productOfNumbers.Add(4)
	productOfNumbers.Add(5)

	expected := []int{20, 60, 120, 240}
	result := []int{
		productOfNumbers.GetProduct(2),
		productOfNumbers.GetProduct(3),
		productOfNumbers.GetProduct(4),
		productOfNumbers.GetProduct(5),
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestProductOfNumbers2(t *testing.T) {
	productOfNumbers := Constructor()
	productOfNumbers.Add(2)
	productOfNumbers.Add(2)
	productOfNumbers.Add(-3)
	productOfNumbers.Add(4)
	productOfNumbers.Add(5)

	expected := []int{20, -60, -120, -240}
	result := []int{
		productOfNumbers.GetProduct(2),
		productOfNumbers.GetProduct(3),
		productOfNumbers.GetProduct(4),
		productOfNumbers.GetProduct(5),
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestProductOfNumbers3(t *testing.T) {
	productOfNumbers := Constructor()
	productOfNumbers.Add(3)
	productOfNumbers.Add(0)
	productOfNumbers.Add(2)
	productOfNumbers.Add(5)
	productOfNumbers.Add(4)

	expected := []int{20, 40, 0, 0}
	result := []int{
		productOfNumbers.GetProduct(2),
		productOfNumbers.GetProduct(3),
		productOfNumbers.GetProduct(4),
		productOfNumbers.GetProduct(5),
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestProductOfNumbers4(t *testing.T) {
	productOfNumbers := Constructor()
	productOfNumbers.Add(3)
	productOfNumbers.Add(0)
	productOfNumbers.Add(2)
	productOfNumbers.Add(5)
	productOfNumbers.Add(0)
	productOfNumbers.Add(4)

	expected := []int{0, 0, 0, 0}
	result := []int{
		productOfNumbers.GetProduct(2),
		productOfNumbers.GetProduct(3),
		productOfNumbers.GetProduct(4),
		productOfNumbers.GetProduct(5),
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
