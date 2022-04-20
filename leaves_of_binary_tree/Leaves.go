package leaves_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	var findLeavesInner func(node *TreeNode) (*TreeNode, []int)
	findLeavesInner = func(node *TreeNode) (*TreeNode, []int) {
		if node == nil {
			return nil, []int{}
		}
		if node.Left == nil && node.Right == nil {
			return node, []int{node.Val}
		}

		var leaves []int
		for _, nextNode := range []*TreeNode{node.Left, node.Right} {
			child, leafValues := findLeavesInner(nextNode)
			if len(leafValues) != 0 {
				leaves = append(leaves, leafValues...)
			}
			if child != nil {
				if child == node.Left {
					node.Left = nil
				} else {
					node.Right = nil
				}
			}
		}
		return nil, leaves
	}

	var result [][]int
	for root.Left != nil || root.Right != nil {
		_, leaves := findLeavesInner(root)
		result = append(result, leaves)
	}
	result = append(result, []int{root.Val})
	return result
}
