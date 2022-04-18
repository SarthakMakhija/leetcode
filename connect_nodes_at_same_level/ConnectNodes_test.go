package connect_nodes_at_same_level

import (
	"reflect"
	"testing"
)

func TestNextRightPointer1(t *testing.T) {
	tree := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}},
	}
	connect(tree)
	values := tree.all()
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestNextRightPointer2(t *testing.T) {
	tree := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4, Left: &Node{Val: 8}, Right: &Node{Val: 9}}, Right: &Node{Val: 5, Left: &Node{Val: 10}, Right: &Node{Val: 11}}},
		Right: &Node{Val: 3, Left: &Node{Val: 6, Left: &Node{Val: 12}, Right: &Node{Val: 13}}, Right: &Node{Val: 7, Left: &Node{Val: 14}, Right: &Node{Val: 15}}},
	}
	connect(tree)
	values := tree.all()
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestNextRightPointer3(t *testing.T) {
	tree := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4, Left: &Node{Val: 8, Left: &Node{Val: 80}, Right: &Node{Val: 81}}, Right: &Node{Val: 9, Left: &Node{Val: 90}, Right: &Node{Val: 91}}}, Right: &Node{Val: 5, Left: &Node{Val: 10, Left: &Node{Val: 100}, Right: &Node{Val: 101}}, Right: &Node{Val: 11, Left: &Node{Val: 110}, Right: &Node{Val: 111}}}},
		Right: &Node{Val: 3, Left: &Node{Val: 6, Left: &Node{Val: 12, Left: &Node{Val: 120}, Right: &Node{Val: 121}}, Right: &Node{Val: 13, Left: &Node{Val: 130}, Right: &Node{Val: 131}}}, Right: &Node{Val: 7, Left: &Node{Val: 14, Left: &Node{Val: 140}, Right: &Node{Val: 141}}, Right: &Node{Val: 15, Left: &Node{Val: 150}, Right: &Node{Val: 151}}}},
	}
	connect(tree)
	values := tree.all()
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 80, 81, 90, 91, 100, 101, 110, 111, 120, 121, 130, 131, 140, 141, 150, 151}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
