package bottom_left_tree_value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var findBottomLeftValueInner func(node *TreeNode, depth int)

	maxDepth, value := -1, -1
	findBottomLeftValueInner = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		findBottomLeftValueInner(node.Left, depth+1)
		if depth > maxDepth {
			maxDepth, value = depth, node.Val
		}
		findBottomLeftValueInner(node.Right, depth+1)
	}
	findBottomLeftValueInner(root, 1)
	return value
}
