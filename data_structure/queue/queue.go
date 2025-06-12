package queue

// Queue 表示一个整数队列结构
// 使用切片（slice）来存储队列元素，实现先进先出（FIFO）的行为
type Queue struct {
	items []int // 存储队列中元素的切片
}

// New 创建并返回一个新的空队列
func New() *Queue {
	q := make([]int, 0, 10) // 队列初始长度为0，容量为10
	return &Queue{items: q}
}

// Enqueue 将一个元素添加到队列尾部
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue 移除并返回队列头部的元素
// 如果队列为空，返回 -1 和 false
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	item := q.items[0]    // 获取队首元素
	q.items = q.items[1:] // 移除队首元素
	return item, true
}

// Peek 返回队列头部的元素但不移除
// 如果队列为空，返回 -1 和 false
func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	return q.items[0], true
}

// IsEmpty 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 返回队列中的元素数量
func (q *Queue) Size() int {
	return len(q.items)
}
