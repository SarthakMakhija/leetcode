package sum_tree

import "testing"

func TestIsSumTree1(t *testing.T) {
	root := &Node{
		data:  10,
		left:  &Node{data: 20, left: &Node{data: 10}, right: &Node{data: 10}},
		right: &Node{data: 30},
	}

	result := isSumTree(root)
	if result != false {
		t.Fatalf("Expected isSumTree to be false but was true")
	}
}

func TestIsSumTree2(t *testing.T) {
	root := &Node{
		data:  3,
		left:  &Node{data: 1},
		right: &Node{data: 2},
	}

	result := isSumTree(root)
	if result != true {
		t.Fatalf("Expected isSumTree to be true but was false")
	}
}
