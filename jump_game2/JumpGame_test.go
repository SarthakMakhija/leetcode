package jump_game2

import "testing"

func TestMinimumJumps1(t *testing.T) {
	minimumJumps := jump([]int{2, 3, 1, 1, 4})
	if minimumJumps != 2 {
		t.Fatalf("Expected %v, received %v", 2, minimumJumps)
	}
}

func TestMinimumJumps2(t *testing.T) {
	minimumJumps := jump([]int{2, 3, 0, 1, 4})
	if minimumJumps != 2 {
		t.Fatalf("Expected %v, received %v", 2, minimumJumps)
	}
}

func TestMinimumJumps3(t *testing.T) {
	minimumJumps := jump([]int{3, 3, 1, 2, 1, 1})
	if minimumJumps != 2 {
		t.Fatalf("Expected %v, received %v", 2, minimumJumps)
	}
}

func TestMinimumJumps4(t *testing.T) {
	minimumJumps := jump([]int{2, 1})
	if minimumJumps != 1 {
		t.Fatalf("Expected %v, received %v", 2, minimumJumps)
	}
}
