package stack

import "testing"

// TestStackPush 测试 Push 方法，确保元素正确入栈。
func TestStackPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)

	if size := s.Size(); size != 2 {
		t.Errorf("期望栈大小为 2，实际为 %d", size)
	}

	if item, ok := s.Peek(); !ok || item != 2 {
		t.Errorf("期望栈顶元素为 2，实际为 %d，ok: %v", item, ok)
	}
}

// TestStackPopEmpty 测试空栈的 Pop 方法，确保返回正确结果。
func TestStackPopEmpty(t *testing.T) {
	s := New()
	item, ok := s.Pop()
	if ok || item != 0 {
		t.Errorf("期望空栈返回 (0, false)，实际为 %d, %v", item, ok)
	}
}

// TestStackIsEmpty 测试 IsEmpty 方法，确保正确判断栈是否为空。
func TestStackIsEmpty(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Error("期望栈为空")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Error("期望栈非空")
	}
}

// TestStackPeek 测试 Peek 方法，确保查看栈顶元素不改变栈状态。
func TestStackPeek(t *testing.T) {
	s := New()
	item, ok := s.Peek()
	if ok || item != 0 {
		t.Errorf("期望空栈返回 (0, false)，实际为 %d, %v", item, ok)
	}

	s.Push(42)
	item, ok = s.Peek()
	if !ok || item != 42 {
		t.Errorf("期望栈顶元素为 42，实际为 %d，ok: %v", item, ok)
	}

	if size := s.Size(); size != 1 {
		t.Errorf("期望 Peek 后栈大小为 1，实际为 %d", size)
	}
}
