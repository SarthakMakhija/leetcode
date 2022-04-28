package find_duplicate_subtrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type nodeValueCount struct {
		nodeValue       int
		leftChildCount  int
		rightChildCount int
	}
	type nodeCount struct {
		node            *TreeNode
		leftChildCount  int
		rightChildCount int
		matched         bool
	}
	treeNodeStructureByValue := make(map[nodeValueCount][]*nodeCount)
	var duplicates []*TreeNode

	type identical bool
	type shouldAdd bool
	isSubTreeAlreadyPresentWith := func(node *TreeNode, leftChildCount int, rightChildCount int) (identical, shouldAdd) {
		existingNodes, ok := treeNodeStructureByValue[nodeValueCount{nodeValue: node.Val, leftChildCount: leftChildCount, rightChildCount: rightChildCount}]
		if !ok {
			return false, false
		}
		var areTreesIdentical func(node *TreeNode, existing *TreeNode, existingLeftChildCount int, existingRightChildCount int) bool
		areTreesIdentical = func(node *TreeNode, existing *TreeNode, existingLeftChildCount int, existingRightChildCount int) bool {
			if node == nil && existing == nil {
				return true
			}
			if (node == nil && existing != nil) || (existing == nil && node != nil) {
				return false
			}
			if node.Val != existing.Val || leftChildCount != existingLeftChildCount || rightChildCount != existingRightChildCount {
				return false
			}
			return areTreesIdentical(node.Left, existing.Left, existingLeftChildCount, existingRightChildCount) &&
				areTreesIdentical(node.Right, existing.Right, existingLeftChildCount, existingRightChildCount)
		}
		for _, existingNodeCount := range existingNodes {
			if areTreesIdentical(node, existingNodeCount.node, existingNodeCount.leftChildCount, existingNodeCount.rightChildCount) {
				if existingNodeCount.matched {
					return true, false
				} else {
					existingNodeCount.matched = true
					return true, true
				}
			}
		}
		return false, false
	}

	var findDuplicates func(node *TreeNode) int
	findDuplicates = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftCount := findDuplicates(node.Left)
		rightCount := findDuplicates(node.Right)
		if areTreesIdentical, shouldBeConsideredDuplicate := isSubTreeAlreadyPresentWith(node, leftCount, rightCount); !areTreesIdentical {
			valueCount := nodeValueCount{nodeValue: node.Val, leftChildCount: leftCount, rightChildCount: rightCount}
			treeNodeStructureByValue[valueCount] = append(
				treeNodeStructureByValue[valueCount], &nodeCount{
					node: node, leftChildCount: leftCount, rightChildCount: rightCount,
				},
			)
		} else {
			if node != root && shouldBeConsideredDuplicate {
				duplicates = append(duplicates, node)
			}
		}
		return leftCount + rightCount + 1
	}
	findDuplicates(root)
	return duplicates
}
