package single_number

import (
	"testing"
)

func TestSingleNumber1(t *testing.T) {
	nums := []int{2, 2, 1}
	number := singleNumber(nums)

	if number != 1 {
		t.Fatalf("Expected %v, received %v", 1, number)
	}
}

func TestSingleNumber2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	number := singleNumber(nums)

	if number != 4 {
		t.Fatalf("Expected %v, received %v", 4, number)
	}
}
