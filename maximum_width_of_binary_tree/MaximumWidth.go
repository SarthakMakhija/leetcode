package maximum_width_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := newNodeTypeQueue()
	if root == nil {
		return 0
	}
	queue.enqueue(root)
	width := 1
	for !queue.isEmpty() {
		size := queue.size()
		childExists := false
		var nilChildCount int32 = 0
		var newSize int32 = 0

		for index := 0; index < size; index++ {
			nodeT := queue.dequeue()
			if nodeT.node.Val == -900 { //signifies nil node
				nilChildCount = nilChildCount + 2*nodeT.count
			} else {
				if nodeT.node.Left != nil {
					if nilChildCount > 0 {
						queue.enqueueNil(nilChildCount)
					}
					queue.enqueue(nodeT.node.Left)
					childExists, nilChildCount, newSize = true, 0, newSize+1+nilChildCount
				} else if nodeT.node.Left == nil && childExists {
					nilChildCount = nilChildCount + 1
				}
				if nodeT.node.Right != nil {
					if nilChildCount > 0 {
						queue.enqueueNil(nilChildCount)
					}
					queue.enqueue(nodeT.node.Right)
					childExists, nilChildCount, newSize = true, 0, newSize+1+nilChildCount
				} else if nodeT.node.Right == nil && childExists {
					nilChildCount = nilChildCount + 1
				}
			}
		}
		if newSize > int32(width) {
			width = int(newSize)
		}
	}
	return width
}

type nodeType struct {
	node  *TreeNode
	count int32
}

type NodeTypeQueue struct {
	items []nodeType
}

func newNodeTypeQueue() *NodeTypeQueue {
	intQueue := &NodeTypeQueue{}
	intQueue.items = []nodeType{}
	return intQueue
}

func (q *NodeTypeQueue) enqueueNil(count int32) {
	q.items = append(q.items, nodeType{node: &TreeNode{Val: -900}, count: count})
}

func (q *NodeTypeQueue) enqueue(node *TreeNode) {
	q.items = append(q.items, nodeType{node: node, count: 1})
}

func (q *NodeTypeQueue) dequeue() nodeType {
	element := q.items[0]
	q.items = q.items[1:len(q.items)]
	return element
}

func (q *NodeTypeQueue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *NodeTypeQueue) size() int {
	return len(q.items)
}
