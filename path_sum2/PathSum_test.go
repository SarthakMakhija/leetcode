package path_sum2

import (
	"reflect"
	"testing"
)

func TestPathSum1(t *testing.T) {
	tree := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
	}
	result := pathSum(tree, 22)
	expected := [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestPathSum2(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	result := pathSum(tree, 5)
	var expected [][]int

	if len(result) != 0 {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
