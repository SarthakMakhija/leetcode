package lowest_common_ancestor_2

import "testing"

func TestLowestCommonAncestor1(t *testing.T) {
	root := &Node{
		Val: 1,
	}
	root.Left = &Node{Val: 2, Parent: root}
	root.Right = &Node{Val: 3, Parent: root}

	lca := lowestCommonAncestor(root.Left, root.Right)
	expected := 1

	if lca.Val != expected {
		t.Fatalf("Expected %v, received %v", expected, lca.Val)
	}
}
