package heap

import (
	"testing"
)

// TestHeapPush 测试 Push 方法，确保值正确插入并保持最小堆性质
func TestHeapPush(t *testing.T) {
	h := New()
	h.Push(3)
	h.Push(1)
	h.Push(4)

	if size := h.Size(); size != 3 {
		t.Errorf("期望堆大小为 3，实际为 %d", size)
	}

	if value, ok := h.Peek(); !ok || value != 1 {
		t.Errorf("期望堆顶值为 1，实际为 %d，ok: %v", value, ok)
	}
}

// TestHeapPop 测试 Pop 方法，确保按最小堆顺序删除并返回最小值
func TestHeapPop(t *testing.T) {
	h := New()
	h.Push(3)
	h.Push(1)
	h.Push(4)

	value, ok := h.Pop()
	if !ok || value != 1 {
		t.Errorf("期望弹出值为 1，实际为 %d，ok: %v", value, ok)
	}

	value, ok = h.Pop()
	if !ok || value != 3 {
		t.Errorf("期望弹出值为 3，实际为 %d，ok: %v", value, ok)
	}

	value, ok = h.Pop()
	if !ok || value != 4 {
		t.Errorf("期望弹出值为 4，实际为 %d，ok: %v", value, ok)
	}

	if size := h.Size(); size != 0 {
		t.Errorf("期望堆大小为 0，实际为 %d", size)
	}
}

// TestHeapPopEmpty 测试空堆的 Pop 方法，确保返回正确结果
func TestHeapPopEmpty(t *testing.T) {
	h := New()
	value, ok := h.Pop()
	if ok || value != -1 {
		t.Errorf("期望空堆 Pop 返回 (-1, false)，实际为 %d，ok: %v", value, ok)
	}
}

// TestHeapPeek 测试 Peek 方法，确保正确查看堆顶值
func TestHeapPeek(t *testing.T) {
	h := New()
	value, ok := h.Peek()
	if ok || value != -1 {
		t.Errorf("期望空堆 Peek 返回 (-1, false)，实际为 %d，ok: %v", value, ok)
	}

	h.Push(2)
	h.Push(1)
	value, ok = h.Peek()
	if !ok || value != 1 {
		t.Errorf("期望堆顶值为 1，实际为 %d，ok: %v", value, ok)
	}

	if size := h.Size(); size != 2 {
		t.Errorf("期望 Peek 后堆大小为 2，实际为 %d", size)
	}
}

// TestHeapIsEmpty 测试 IsEmpty 方法，确保正确判断堆是否为空
func TestHeapIsEmpty(t *testing.T) {
	h := New()
	if !h.IsEmpty() {
		t.Error("期望堆为空")
	}

	h.Push(1)
	if h.IsEmpty() {
		t.Error("期望堆非空")
	}
}

// TestHeapSize 测试 Size 方法，确保返回正确的元素数量
func TestHeapSize(t *testing.T) {
	h := New()
	if size := h.Size(); size != 0 {
		t.Errorf("期望堆大小为 0，实际为 %d", size)
	}

	h.Push(1)
	h.Push(2)
	if size := h.Size(); size != 2 {
		t.Errorf("期望堆大小为 2，实际为 %d", size)
	}
}

// TestHeapProperty 测试最小堆性质，确保堆顶始终是最小值
func TestHeapProperty(t *testing.T) {
	h := New()
	values := []int{5, 2, 8, 1, 9, 3}
	for _, v := range values {
		h.Push(v)
	}

	expected := []int{1, 2, 3, 5, 8, 9} // 排序后的值
	for i, exp := range expected {
		value, ok := h.Pop()
		if !ok || value != exp {
			t.Errorf("第 %d 次弹出期望值为 %d，实际为 %d，ok: %v", i+1, exp, value, ok)
		}
	}
}
