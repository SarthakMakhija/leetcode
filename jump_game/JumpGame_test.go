package jump_game

import "testing"

func TestCanJump1(t *testing.T) {
	canJump := canJump([]int{2, 3, 1, 1, 4})
	if canJump != true {
		t.Fatalf("Expected jump to be possible but was not")
	}
}

func TestCanJump2(t *testing.T) {
	canJump := canJump([]int{3, 2, 1, 0, 4})
	if canJump != false {
		t.Fatalf("Expected jump to be impossible but was possible")
	}
}

func TestCanJump3(t *testing.T) {
	canJump := canJump([]int{2, 1, 0, 1, 4})
	if canJump != false {
		t.Fatalf("Expected jump to be impossible but was possible")
	}
}
