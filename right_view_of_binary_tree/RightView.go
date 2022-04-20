package right_view_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var rightView []int
	maxLevel := 0

	var rightViewInner func(node *TreeNode, level int)
	rightViewInner = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxLevel {
			rightView = append(rightView, node.Val)
			maxLevel = level
		}
		rightViewInner(node.Right, level+1)
		rightViewInner(node.Left, level+1)
	}
	rightViewInner(root, 1)
	return rightView
}
