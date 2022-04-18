package sum_tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func isSumTree(root *Node) bool {
	var isSumTreeInner func(node *Node) (int, bool)
	isSumTreeInner = func(node *Node) (int, bool) {
		switch {
		case node == nil:
			return 0, true
		case node.left == nil && node.right == nil:
			return node.data, true
		}
		leftTreeSum, leftTreeSumOk := isSumTreeInner(node.left)
		if leftTreeSumOk {
			rightTreeSum, rightTreeSumOk := isSumTreeInner(node.right)
			if rightTreeSumOk {
				if node.data == leftTreeSum+rightTreeSum {
					return node.data + leftTreeSum + rightTreeSum, true
				} else {
					return -1, false
				}
			}
		}
		return leftTreeSum, leftTreeSumOk
	}
	_, ok := isSumTreeInner(root)
	return ok
}
