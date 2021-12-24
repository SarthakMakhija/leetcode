package search_2D_matrix

import "testing"

func TestSearch2DMatrix1(t *testing.T) {
	target := 3
	found := searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, target)

	if found != true {
		t.Fatalf("Expected %v to be found but was not", target)
	}
}

func TestSearch2DMatrix2(t *testing.T) {
	target := 0
	found := searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, target)

	if found != false {
		t.Fatalf("Expected %v to not be found but was", target)
	}
}

func TestSearch2DMatrix3(t *testing.T) {
	target := 34
	found := searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, target)

	if found != true {
		t.Fatalf("Expected %v to be found but was not", target)
	}
}

func TestSearch2DMatrix4(t *testing.T) {
	target := 60
	found := searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, target)

	if found != true {
		t.Fatalf("Expected %v to be found but was not", target)
	}
}

func TestSearch2DMatrix5(t *testing.T) {
	target := 61
	found := searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, target)

	if found != false {
		t.Fatalf("Expected %v to not be found but was", target)
	}
}
