package stack

import "testing"

// TestStack_Push 测试 Push 方法，确保元素正确入栈。
func TestStack_Push(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if size := s.Size(); size != 2 {
		t.Errorf("期望栈大小为 2，实际为 %d", size)
	}

	if item, ok := s.Peek(); !ok || item != 2 {
		t.Errorf("期望栈顶元素为 2，实际为 %d，ok: %v", item, ok)
	}
}

// TestStack_PopEmpty 测试空栈的 Pop 方法，确保返回正确结果。
func TestStack_PopEmpty(t *testing.T) {
	s := NewStack()
	item, ok := s.Pop()
	if ok || item != 0 {
		t.Errorf("期望空栈返回 (0, false)，实际为 %d, %v", item, ok)
	}
}

// TestStack_IsEmpty 测试 IsEmpty 方法，确保正确判断栈是否为空。
func TestStack_IsEmpty(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() {
		t.Error("期望栈为空")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Error("期望栈非空")
	}
}

// TestStack_Peek 测试 Peek 方法，确保查看栈顶元素不改变栈状态。
func TestStack_Peek(t *testing.T) {
	s := NewStack()
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
