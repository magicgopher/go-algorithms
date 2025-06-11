package queue

// Queue 表示一个先进先出（FIFO）的整数队列，基于切片实现。
type Queue struct {
	items []int
}

// New 创建并返回一个新的空队列。
func New() *Queue {
	return &Queue{items: []int{}}
}

// Enqueue 将一个元素添加到队列尾部。
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue 移除并返回队列头部的元素。
// 如果队列为空，返回 -1 和 false。
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front 返回队列头部的元素但不移除。
// 如果队列为空，返回 -1 和 false。
func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	return q.items[0], true
}

// IsEmpty 检查队列是否为空。
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 返回队列中的元素数量。
func (q *Queue) Size() int {
	return len(q.items)
}
