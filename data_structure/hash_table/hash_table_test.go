package hash_table

import "testing"

// TestPutAndGet 测试插入和查找操作
func TestPutAndGet(t *testing.T) {
	ht := New(10)

	err := ht.Put("name", "Alice")
	if err != nil {
		t.Errorf("Put 失败: %v", err)
	}

	val, ok := ht.Get("name")
	if !ok || val != "Alice" {
		t.Errorf("Get 失败，预期 Alice，实际 %s", val)
	}

	// 测试更新值
	ht.Put("name", "Bob")
	val, _ = ht.Get("name")
	if val != "Bob" {
		t.Errorf("更新失败，预期 Bob，实际 %s", val)
	}
}

// TestRemove 测试删除操作
func TestRemove(t *testing.T) {
	ht := New(10)
	ht.Put("age", "30")

	removed := ht.Remove("age")
	if !removed {
		t.Errorf("Remove 失败，未能删除已有键")
	}

	_, ok := ht.Get("age")
	if ok {
		t.Errorf("Get 应该失败，键 age 已被删除")
	}
}

// TestCollision 测试哈希冲突处理
func TestCollision(t *testing.T) {
	ht := New(5)

	// 人为制造冲突（使用相同的哈希函数）
	ht.Put("a", "value1")
	ht.Put("f", "value2") // 假设两个键冲突

	if val, ok := ht.Get("a"); !ok || val != "value1" {
		t.Errorf("冲突键 a 查找失败")
	}
	if val, ok := ht.Get("f"); !ok || val != "value2" {
		t.Errorf("冲突键 f 查找失败")
	}
}

// TestTableFull 测试表满时插入失败
func TestTableFull(t *testing.T) {
	ht := New(2)

	ht.Put("a", "1")
	ht.Put("b", "2")
	err := ht.Put("c", "3")

	if err == nil {
		t.Errorf("应在表满时返回错误，但未返回")
	}
}
