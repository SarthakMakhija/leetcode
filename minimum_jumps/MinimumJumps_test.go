package minimum_jumps

import (
	"testing"
)

func TestMinJumps1(t *testing.T) {
	jumps := jump([]int{2, 3, 1, 1, 4})
	expected := 2

	if jumps != expected {
		t.Fatalf("Expected jumps to be %v, received %v", jumps, expected)
	}
}

func TestMinJumps2(t *testing.T) {
	jumps := jump([]int{2, 3, 0, 1, 4})
	expected := 2

	if jumps != expected {
		t.Fatalf("Expected jumps to be %v, received %v", jumps, expected)
	}
}

func TestMinJumps3(t *testing.T) {
	jumps := jump([]int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9})
	expected := 3

	if jumps != expected {
		t.Fatalf("Expected jumps to be %v, received %v", jumps, expected)
	}
}
