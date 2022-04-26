package nodes_at_k_distance

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var findNodesAtKWithAncestor func(
		node *TreeNode,
		remainingK int,
		ancestors []*TreeNode,
		foundTarget bool,
	) []int
	var findNodesAtDistanceFromAncestor func(node *TreeNode, skipNode *TreeNode, k int) []int

	isATargetNodeAncestor := func(node *TreeNode, allAncestors []*TreeNode) bool {
		for _, ancestor := range allAncestors {
			if ancestor == node {
				return true
			}
		}
		return false
	}

	var targetNodeAncestors []*TreeNode
	findNodesAtKWithAncestor = func(
		node *TreeNode,
		remainingK int,
		ancestors []*TreeNode,
		foundTarget bool,
	) []int {
		if node == nil {
			return []int{}
		}
		if remainingK == 0 && foundTarget {
			return []int{node.Val}
		}
		var remaining = remainingK
		var found = foundTarget
		if node.Val == target.Val {
			targetNodeAncestors = ancestors
			found = true
			if remainingK == 0 && found {
				return []int{node.Val}
			}
		}
		if found {
			remaining = remaining - 1
		}
		return append(
			findNodesAtKWithAncestor(node.Left, remaining, append(ancestors, node), found),
			findNodesAtKWithAncestor(node.Right, remaining, append(ancestors, node), found)...)
	}

	findNodesAtDistanceFromAncestor = func(node *TreeNode, skipNode *TreeNode, k int) []int {
		if node == nil {
			return []int{}
		}
		if node.Val == skipNode.Val {
			return []int{}
		}
		if k == 0 {
			if node.Val != skipNode.Val && !isATargetNodeAncestor(node, targetNodeAncestors) {
				return []int{node.Val}
			}
			return []int{}
		}
		return append(
			findNodesAtDistanceFromAncestor(node.Left, skipNode, k-1),
			findNodesAtDistanceFromAncestor(node.Right, skipNode, k-1)...)
	}

	nodesAtKDistance := findNodesAtKWithAncestor(root, k, []*TreeNode{}, false)
	reduceKBy := 1
	var previousAncestor *TreeNode
	for index := len(targetNodeAncestors) - 1; index >= 0; index-- {
		ancestor := targetNodeAncestors[index]
		k = k - reduceKBy
		switch {
		case k < 0:
			break
		case k == 0:
			nodesAtKDistance = append(nodesAtKDistance, ancestor.Val)
			continue
		}
		var skipNode *TreeNode
		if previousAncestor == nil {
			skipNode = target
		} else {
			skipNode = previousAncestor
		}
		nodesAtKDistance = append(nodesAtKDistance, findNodesAtDistanceFromAncestor(ancestor, skipNode, k)...)
		previousAncestor = ancestor
	}
	return nodesAtKDistance
}
