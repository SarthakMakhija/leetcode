package step_by_step_direction

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	type nodeLevel struct {
		node  *TreeNode
		level int
	}

	reverse := func(nodeLevels []nodeLevel) []nodeLevel {
		for i := 0; i < len(nodeLevels)/2; i++ {
			j := len(nodeLevels) - i - 1
			nodeLevels[i], nodeLevels[j] = nodeLevels[j], nodeLevels[i]
		}
		return nodeLevels
	}

	var findPath func(
		node *TreeNode,
		startValue,
		destinationValue,
		level int,
		stAncestors []nodeLevel,
		destAncestors []nodeLevel,
	)

	var startAncestors, destinationAncestors []nodeLevel
	startFound, destFound := false, false

	findPath = func(
		node *TreeNode,
		startValue,
		destinationValue,
		level int,
		stAncestors []nodeLevel,
		destAncestors []nodeLevel,
	) {
		if node == nil {
			return
		}
		if node.Val == startValue {
			startAncestors = append(stAncestors, nodeLevel{node: node, level: level})
			startFound = true
		}
		if node.Val == destinationValue {
			destinationAncestors = append(destAncestors, nodeLevel{node: node, level: level})
			destFound = true
		}
		if startFound && destFound {
			return
		}

		newStartAncestors, newDestinationAncestors := stAncestors, destAncestors
		if !startFound {
			newStartAncestors = append(stAncestors, nodeLevel{node: node, level: level})
		}
		if !destFound {
			newDestinationAncestors = append(destAncestors, nodeLevel{node: node, level: level})
		}

		findPath(
			node.Left,
			startValue,
			destinationValue,
			level+1,
			newStartAncestors,
			newDestinationAncestors,
		)
		if startFound && destFound {
			return
		}
		if !startFound {
			newStartAncestors = append(stAncestors, nodeLevel{node: node, level: level})
		}
		if !destFound {
			newDestinationAncestors = append(destAncestors, nodeLevel{node: node, level: level})
		}
		findPath(
			node.Right,
			startValue,
			destinationValue,
			level+1,
			newStartAncestors,
			newDestinationAncestors,
		)
	}

	findPath(root, startValue, destValue, 1, []nodeLevel{}, []nodeLevel{})
	startValuePath, destinationValuePath := reverse(startAncestors), reverse(destinationAncestors)
	aLength, bLength := len(startValuePath), len(destinationValuePath)

	var totalCommon int
	for aIndex, bIndex := aLength-1, bLength-1; aIndex >= 0 && bIndex >= 0; aIndex, bIndex = aIndex-1, bIndex-1 {
		if startValuePath[aIndex].node.Val == destinationValuePath[bIndex].node.Val {
			totalCommon = totalCommon + 1
		} else {
			break
		}
	}

	var path strings.Builder
	for index := 0; index < aLength-1-(totalCommon-1); index++ {
		if startValuePath[index].level > startValuePath[index+1].level {
			path.WriteRune('U')
		} else if startValuePath[index].node == startValuePath[index+1].node.Left {
			path.WriteRune('L')
		} else if startValuePath[index].node == startValuePath[index+1].node.Right {
			path.WriteRune('R')
		}
	}
	for index := bLength - 1 - (totalCommon - 1); index > 0; index-- {
		if destinationValuePath[index-1].level < destinationValuePath[index].level {
			path.WriteRune('U')
		} else if destinationValuePath[index-1].node == destinationValuePath[index].node.Left {
			path.WriteRune('L')
		} else if destinationValuePath[index-1].node == destinationValuePath[index].node.Right {
			path.WriteRune('R')
		}
	}
	return path.String()
}
