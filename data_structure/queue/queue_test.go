package queue

import (
	"testing"
)

// TestQueueEnqueue 测试 Enqueue 方法，确保元素正确入队。
func TestQueueEnqueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)

	if size := q.Size(); size != 2 {
		t.Errorf("期望队列大小为 2，实际为 %d", size)
	}

	if item, ok := q.Front(); !ok || item != 1 {
		t.Errorf("期望队列头部元素为 1，实际为 %d，ok: %v", item, ok)
	}
}

// TestQueueDequeue 测试 Dequeue 方法，确保元素按先进先出顺序出队。
func TestQueueDequeue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)

	item, ok := q.Dequeue()
	if !ok || item != 1 {
		t.Errorf("期望出队元素为 1，实际为 %d，ok: %v", item, ok)
	}

	item, ok = q.Dequeue()
	if !ok || item != 2 {
		t.Errorf("期望出队元素为 2，实际为 %d，ok: %v", item, ok)
	}

	if size := q.Size(); size != 0 {
		t.Errorf("期望队列大小为 0，实际为 %d", size)
	}
}

// TestQueueDequeueEmpty 测试空队列的 Dequeue 方法，确保返回正确结果。
func TestQueueDequeueEmpty(t *testing.T) {
	q := New()
	item, ok := q.Dequeue()
	if ok || item != -1 {
		t.Errorf("期望空队列 Dequeue 返回 (-1, false)，实际为 %d，ok: %v", item, ok)
	}
}

// TestQueueFront 测试 Front 方法，确保正确查看头部元素。
func TestQueueFront(t *testing.T) {
	q := New()
	item, ok := q.Front()
	if ok || item != -1 {
		t.Errorf("期望空队列 Front 返回 (-1, false)，实际为 %d，ok: %v", item, ok)
	}

	q.Enqueue(42)
	item, ok = q.Front()
	if !ok || item != 42 {
		t.Errorf("期望队列头部元素为 42，实际为 %d，ok: %v", item, ok)
	}

	if size := q.Size(); size != 1 {
		t.Errorf("期望 Front 后队列大小为 1，实际为 %d", size)
	}
}

// TestQueueIsEmpty 测试 IsEmpty 方法，确保正确判断队列是否为空。
func TestQueueIsEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Error("期望队列为空")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("期望队列非空")
	}
}

// TestQueueSize 测试 Size 方法，确保返回正确的元素数量。
func TestQueueSize(t *testing.T) {
	q := New()
	if size := q.Size(); size != 0 {
		t.Errorf("期望队列大小为 0，实际为 %d", size)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	if size := q.Size(); size != 2 {
		t.Errorf("期望队列大小为 2，实际为 %d", size)
	}
}
