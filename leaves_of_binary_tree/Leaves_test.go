package leaves_of_binary_tree

import (
	"reflect"
	"testing"
)

func TestFindLeaves1(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	leaves := findLeaves(root)
	expected := [][]int{
		{4, 5, 3},
		{2},
		{1},
	}

	if !reflect.DeepEqual(expected, leaves) {
		t.Fatalf("Expected %v, received %v", expected, leaves)
	}
}
