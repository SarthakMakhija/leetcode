package range_sum_query_mutable

import (
	"testing"
)

func TestRangeSumQuery1(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})
	sum := numArray.SumRange(0, 2)
	expected := 9

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestRangeSumQuery2(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})

	numArray.Update(0, 10)
	sum := numArray.SumRange(0, 2)
	expected := 18

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestRangeSumQuery3(t *testing.T) {
	numArray := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8})
	sum := numArray.SumRange(2, 6)
	expected := 25

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestRangeSumQuery4(t *testing.T) {
	numArray := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8})

	numArray.Update(2, 30)
	sum := numArray.SumRange(2, 6)
	expected := 52

	if expected != sum {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}
