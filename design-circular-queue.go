
// MyCircularQueue 结构体
type MyCircularQueue struct {
	queue []int
	size  int
	head  int
	tail  int
}

// Constructor initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		size:  0,
		head:  0,
		tail:  -1,
	}
}

// EnQueue insert an element into the circular queue. Return true if the operation is successful.
func (m *MyCircularQueue) EnQueue(value int) bool {
	if m.IsFull() {
		return false
	}
	m.tail++
	if m.tail >= len(m.queue) {
		m.tail = 0
	}
	m.queue[m.tail] = value
	m.size++
	return true
}

// DeQueue delete an element from the circular queue. Return true if the operation is successful.
func (m *MyCircularQueue) DeQueue() bool {
	if m.IsEmpty() {
		return false
	}
	m.head++
	if m.head >= len(m.queue) {
		m.head = 0
	}
	m.size--
	return true
}

// Front get the front item from the queue.
func (m *MyCircularQueue) Front() int {
	if m.IsEmpty() {
		return -1
	}
	return m.queue[m.head]
}

// Rear get the last item from the queue. */
func (m *MyCircularQueue) Rear() int {
	if m.IsEmpty() {
		return -1
	}
	return m.queue[m.tail]
}

// IsEmpty checks whether the circular queue is empty or not. */
func (m *MyCircularQueue) IsEmpty() bool {
	return m.size == 0
}

// IsFull checks whether the circular queue is full or not. */
func (m *MyCircularQueue) IsFull() bool {
	return m.size == len(m.queue)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
