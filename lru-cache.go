// package problem0146

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// doublyLinkedNode 是双向链节点
type doublyLinkedNode struct {
	key, val   int
	next, prev *doublyLinkedNode
}

// LRUCache 利用 双向链条 + hashtabl 实现
type LRUCache struct {
	nodes         map[int]*doublyLinkedNode
	capacity, len int
	first, last   *doublyLinkedNode
}

// Constructor 创建容量为 capacity 的 cache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		nodes:    make(map[int]*doublyLinkedNode, capacity),
		capacity: capacity,
	}
}

// Get 获取 cache 中的数据
func (c *LRUCache) Get(key int) int {
	node, ok := c.nodes[key]
	if ok {
		c.moveToFirst(node)
		return node.val
	}
	return -1
}

// Put is 放入新数据
func (c *LRUCache) Put(key int, value int) {
	node, ok := c.nodes[key]
	if ok {
		node.val = value
		c.moveToFirst(node)
	} else {
		if c.capacity == c.len {
			delete(c.nodes, c.last.key)
			c.removeLast()
		} else {
			c.len++
		}
		l := &doublyLinkedNode{
			key: key,
			val: value,
		}
		c.nodes[key] = l
		c.insertToFirst(l)
	}
}

func (c *LRUCache) removeLast() {
	//如果是单节点, 删除这个节点. 删除 first 引用
	//如果不是单节点, 删除最后一个节点. 然后将上一个节点变成最后一个节点
	if c.last.prev != nil {
		//将上一个节点变成最后一个节点
		c.last.prev.next = nil
	} else {
		c.first = nil
	}
	c.last = c.last.prev
}

func (c *LRUCache) moveToFirst(node *doublyLinkedNode) {
	//将已经存在的节点 放到第一个
	//是否是单节点
	//是否是最后节点
	switch node {
	case c.first:
		return
	case c.last:
		c.removeLast()
	default:
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	c.insertToFirst(node)
}

func (c *LRUCache) insertToFirst(node *doublyLinkedNode) {
	//如单节点, 这个节点要变成 last 节点
	if c.last == nil {
		c.last = node
	} else {
		node.next = c.first
		c.first.prev = node
	}

	c.first = node

}
