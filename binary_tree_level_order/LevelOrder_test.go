package binary_tree_level_order

import (
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	levelOrderValues := levelOrder(tree)
	expected := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}

	if !reflect.DeepEqual(expected, levelOrderValues) {
		t.Fatalf("Expected %v, received %v", expected, levelOrderValues)
	}
}
