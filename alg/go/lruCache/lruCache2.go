package lruCache

type ListNode struct {
	key, val   int
	prev, next *ListNode
}

type LRUCache2 struct {
	capacity   int
	size       int
	cache      map[int]*ListNode
	head, tail *ListNode
}

func Constructor2(capacity int) LRUCache2 {
	head, tail := &ListNode{}, &ListNode{}
	head.next = tail
	tail.prev = head
	return LRUCache2{capacity: capacity, cache: make(map[int]*ListNode), head: head, tail: tail}
}

func (c *LRUCache2) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		return node.val
	}
	return -1
}

func (c *LRUCache2) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		node.val = value
		c.moveToHead(node)
		return
	}

	node := &ListNode{key: key, val: value}
	c.cache[key] = node
	c.addToHead(node)
	c.size++

	if c.size > c.capacity {
		removed := c.removeTail()
		delete(c.cache, removed.key)
		c.size--
	}
}

func (c *LRUCache2) moveToHead(node *ListNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache2) removeNode(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache2) addToHead(node *ListNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache2) removeTail() *ListNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
