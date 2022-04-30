package lowest_common_ancestor_2

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	parents := make(map[int]bool)
	var lowestCommonAncestorInner func(first, second *Node) *Node
	lowestCommonAncestorInner = func(first, second *Node) *Node {
		if first == second {
			return first
		}
		if first != nil {
			if _, ok := parents[first.Val]; ok {
				return first
			}
			parents[first.Val] = true
		}
		if second != nil {
			if _, ok := parents[second.Val]; ok {
				return second
			}
			parents[second.Val] = true
		}
		var firstParent, secondParent *Node
		if first != nil {
			firstParent = first.Parent
		}
		if second != nil {
			secondParent = second.Parent
		}
		return lowestCommonAncestorInner(firstParent, secondParent)
	}
	return lowestCommonAncestorInner(p, q)
}
