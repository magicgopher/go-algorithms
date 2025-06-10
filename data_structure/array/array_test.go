package array

import "testing"

// TestArrayAppend 测试 Append 方法，确保元素正确追加到数组。
func TestArrayAppend(t *testing.T) {
	a := New()
	a.Append(1)
	a.Append(2)

	if size := a.Size(); size != 2 {
		t.Errorf("期望数组大小为 2，实际为 %d", size)
	}

	if item, ok := a.Get(0); !ok || item != 1 {
		t.Errorf("期望索引 0 处元素为 1，实际为 %d，ok: %v", item, ok)
	}

	if item, ok := a.Get(1); !ok || item != 2 {
		t.Errorf("期望索引 1 处元素为 2，实际为 %d，ok: %v", item, ok)
	}
}

// TestArrayGet 测试 Get 方法，确保正确获取元素。
func TestArrayGet(t *testing.T) {
	// 测试空数组
	a := New()
	if item, ok := a.Get(0); ok || item != -1 {
		t.Errorf("空数组 Get(0) 期望返回 (0, false)，实际为 %d，ok: %v", item, ok)
	}

	// 添加元素 [10, 20, 30]
	a.Append(10)
	a.Append(20)
	a.Append(30)

	// 测试有效索引
	tests := []struct {
		index    int
		expected int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
	}
	for _, tt := range tests {
		if item, ok := a.Get(tt.index); !ok || item != tt.expected {
			t.Errorf("Get(%d) 期望返回 (%d, true)，实际为 %d，ok: %v", tt.index, tt.expected, item, ok)
		}
	}

	// 测试无效索引
	invalidTests := []int{-1, 3, 100}
	for _, index := range invalidTests {
		if item, ok := a.Get(index); ok || item != -1 {
			t.Errorf("Get(%d) 期望返回 (0, false)，实际为 %d，ok: %v", index, item, ok)
		}
	}
}

// TestArraySet 测试 Set 方法，确保正确设置元素。
func TestArraySet(t *testing.T) {
	a := New()
	a.Append(1)
	a.Append(2)

	if !a.Set(1, 42) {
		t.Error("期望 Set 方法成功，实际失败")
	}

	if item, ok := a.Get(1); !ok || item != 42 {
		t.Errorf("期望索引 1 处元素为 42，实际为 %d，ok: %v", item, ok)
	}

	if a.Set(2, 99) {
		t.Error("期望无效索引 Set 方法失败，实际成功")
	}
}

// TestArrayRemove 测试 Remove 方法，确保正确删除元素。
func TestArrayRemove(t *testing.T) {
	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(3)

	item, ok := a.Remove(1)
	if !ok || item != 2 {
		t.Errorf("期望删除索引 1 处元素为 2，实际为 %d，ok: %v", item, ok)
	}

	if size := a.Size(); size != 2 {
		t.Errorf("期望数组大小为 2，实际为 %d", size)
	}

	if item, ok := a.Get(1); !ok || item != 3 {
		t.Errorf("期望索引 1 处元素为 3，实际为 %d，ok: %v", item, ok)
	}

	item, ok = a.Remove(5)
	if ok || item != 0 {
		t.Errorf("期望无效索引删除返回 (0, false)，实际为 %d，ok: %v", item, ok)
	}
}

// TestArrayIsEmpty 测试 IsEmpty 方法，确保正确判断数组是否为空。
func TestArrayIsEmpty(t *testing.T) {
	a := New()
	if !a.IsEmpty() {
		t.Error("期望数组为空")
	}

	a.Append(1)
	if a.IsEmpty() {
		t.Error("期望数组非空")
	}
}

// TestArraySize 测试 Size 方法，确保返回正确的元素数量。
func TestArraySize(t *testing.T) {
	a := New()
	if size := a.Size(); size != 0 {
		t.Errorf("期望数组大小为 0，实际为 %d", size)
	}

	a.Append(1)
	a.Append(2)
	if size := a.Size(); size != 2 {
		t.Errorf("期望数组大小为 2，实际为 %d", size)
	}
}
