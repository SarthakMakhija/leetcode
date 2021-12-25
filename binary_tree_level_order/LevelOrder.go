package binary_tree_level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeByLevel struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var nodes []*NodeByLevel
	nodes = append(nodes, &NodeByLevel{node: root, level: 0})

	var output [][]int
	level := 0
	for level != len(nodes) {
		nodesByLevel, values := findAllNodesBy(nodes, level)
		if len(values) != 0 {
			output = append(output, values)
		}
		for _, levelNode := range nodesByLevel {
			if levelNode.node.Left != nil {
				nodes = append(nodes, &NodeByLevel{node: levelNode.node.Left, level: level + 1})
			}
			if levelNode.node.Right != nil {
				nodes = append(nodes, &NodeByLevel{node: levelNode.node.Right, level: level + 1})
			}
		}
		level = level + 1
	}
	return output
}

func findAllNodesBy(nodes []*NodeByLevel, level int) ([]*NodeByLevel, []int) {
	var levelNodes []*NodeByLevel
	var values []int
	for _, levelNode := range nodes {
		if levelNode.level == level {
			levelNodes = append(levelNodes, levelNode)
			values = append(values, levelNode.node.Val)
		}
	}
	return levelNodes, values
}
