package house_robber

import "testing"

func TestRob1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	robberyAmount := rob(nums)
	expected := 4

	if expected != robberyAmount {
		t.Fatalf("Expected %v, received %v", expected, robberyAmount)
	}
}

func TestRob2(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	robberyAmount := rob(nums)
	expected := 12

	if expected != robberyAmount {
		t.Fatalf("Expected %v, received %v", expected, robberyAmount)
	}
}

func TestRob3(t *testing.T) {
	var nums []int
	robberyAmount := rob(nums)
	expected := 0

	if expected != robberyAmount {
		t.Fatalf("Expected %v, received %v", expected, robberyAmount)
	}
}

func TestRob4(t *testing.T) {
	nums := []int{2, 7, 1, 9, 8}
	robberyAmount := rob(nums)
	expected := 16

	if expected != robberyAmount {
		t.Fatalf("Expected %v, received %v", expected, robberyAmount)
	}
}

func TestRob5(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	robberyAmount := rob(nums)
	expected := 4

	if expected != robberyAmount {
		t.Fatalf("Expected %v, received %v", expected, robberyAmount)
	}
}
