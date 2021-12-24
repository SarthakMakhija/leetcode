package lru_cache

type LRUCache struct {
	capacity  int
	nodeByKey map[int]*DoublyListNode
	listHead  *DoublyListNode
	listTail  *DoublyListNode
}

type DoublyListNode struct {
	key, value int
	next       *DoublyListNode
	previous   *DoublyListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		nodeByKey: make(map[int]*DoublyListNode),
	}
}

func (lruCache *LRUCache) Get(key int) int {
	node := lruCache.get(key)
	if node == nil {
		return -1
	}
	return node.value
}

func (lruCache *LRUCache) get(key int) *DoublyListNode {
	node, ok := lruCache.nodeByKey[key]
	if !ok {
		return nil
	}
	if node == lruCache.listHead {
		return node
	}
	if node == lruCache.listTail {
		lruCache.listTail = node.previous
	}
	node.previous.next = node.next
	if node.next != nil {
		node.next.previous = node.previous
	}
	node.previous = nil
	node.next = nil
	lruCache.prepend(node)
	return node
}

func (lruCache *LRUCache) Put(key int, value int) {
	node := lruCache.get(key)
	if node != nil {
		node.value = value
		return
	}
	if len(lruCache.nodeByKey) >= lruCache.capacity {
		lruCache.evict()
	}
	lruCache.nodeByKey[key] = lruCache.put(key, value)
}

func (lruCache *LRUCache) put(key, value int) *DoublyListNode {
	newNode := &DoublyListNode{key: key, value: value}
	if lruCache.listHead == nil {
		lruCache.listHead, lruCache.listTail = newNode, newNode
	} else {
		lruCache.prepend(newNode)
	}
	return newNode
}

func (lruCache *LRUCache) prepend(node *DoublyListNode) {
	node.next = lruCache.listHead
	lruCache.listHead.previous = node
	lruCache.listHead = node
}

func (lruCache *LRUCache) evict() {
	if lruCache.listHead == lruCache.listTail {
		lruCache.listHead = nil
	}
	if lruCache.listTail != nil {
		delete(lruCache.nodeByKey, lruCache.listTail.key)
		previous := lruCache.listTail.previous
		if previous != nil {
			previous.next = nil
		}
		lruCache.listTail.previous = nil
		lruCache.listTail = previous
	}
}
