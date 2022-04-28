package find_duplicate_subtrees

import (
	"reflect"
	"testing"
)

func TestFindDuplicateSubtrees1(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 4}},
	}

	var result []int
	for _, node := range findDuplicateSubtrees(root) {
		result = append(result, node.Val)
	}
	expectedRootValues := []int{4, 2}
	if !reflect.DeepEqual(expectedRootValues, result) {
		t.Fatalf("Expected %v, received %v", expectedRootValues, result)
	}
}

func TestFindDuplicateSubtrees2(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 1},
	}

	var result []int
	for _, node := range findDuplicateSubtrees(root) {
		result = append(result, node.Val)
	}
	expectedRootValues := []int{1}
	if !reflect.DeepEqual(expectedRootValues, result) {
		t.Fatalf("Expected %v, received %v", expectedRootValues, result)
	}
}

func TestFindDuplicateSubtrees3(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}

	var result []int
	for _, node := range findDuplicateSubtrees(root) {
		result = append(result, node.Val)
	}
	expectedRootValues := []int{3, 2}
	if !reflect.DeepEqual(expectedRootValues, result) {
		t.Fatalf("Expected %v, received %v", expectedRootValues, result)
	}
}
