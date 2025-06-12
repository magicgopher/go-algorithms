package queue

import (
	"testing"
)

// TestEnqueue 测试 Enqueue 方法是否正确添加元素到队列中
func TestEnqueue(t *testing.T) {
	q := New()
	q.Enqueue(10)
	q.Enqueue(20)

	if q.Size() != 2 {
		t.Errorf("预期队列大小为 2，实际为 %d", q.Size())
	}

	if val, _ := q.Peek(); val != 10 {
		t.Errorf("预期队首元素为 10，实际为 %d", val)
	}
}

// TestDequeue 测试 Dequeue 方法是否正确移除并返回队首元素
func TestDequeue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)

	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("预期出队元素为 1，实际为 %d", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("预期出队元素为 2，实际为 %d", val)
	}

	// 测试空队列出队
	val, ok = q.Dequeue()
	if ok {
		t.Errorf("预期空队列返回 false，实际返回 true，值为 %d", val)
	}
}

// TestPeek 测试 Peek 方法是否正确返回队首元素而不移除
func TestPeek(t *testing.T) {
	q := New()
	q.Enqueue(100)

	val, ok := q.Peek()
	if !ok || val != 100 {
		t.Errorf("Peek 失败，预期为 100，实际为 %d", val)
	}

	// Peek 不应移除元素
	if q.Size() != 1 {
		t.Errorf("Peek 后队列大小应仍为 1，实际为 %d", q.Size())
	}
}

// TestIsEmpty 测试队列是否能正确判断为空状态
func TestIsEmpty(t *testing.T) {
	q := New()

	if !q.IsEmpty() {
		t.Errorf("新建队列应为空")
	}

	q.Enqueue(5)

	if q.IsEmpty() {
		t.Errorf("添加元素后队列不应为空")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("移除所有元素后队列应为空")
	}
}
