package linked_list

import "testing"

// TestLinkedList 测试链表的各种操作
func TestLinkedList(t *testing.T) {
	// 创建一个新的空链表
	ll := NewLinkedList()

	// 测试空链表的初始状态
	if ll.Size != 0 {
		t.Errorf("预期链表大小为0，实际为 %d", ll.Size)
	}
	if ll.Contains(1) {
		t.Error("空链表不应包含值1")
	}

	// 测试插入操作
	ll.Insert(1) // 插入第一个节点
	ll.Insert(2) // 插入第二个节点
	ll.Insert(3) // 插入第三个节点

	if ll.Size != 3 {
		t.Errorf("插入3个节点后，预期链表大小为3，实际为 %d", ll.Size)
	}

	// 测试包含检查
	if !ll.Contains(2) {
		t.Error("链表应包含值2")
	}
	if ll.Contains(4) {
		t.Error("链表不应包含值4")
	}

	// 测试删除操作
	if !ll.Delete(2) {
		t.Error("删除值2失败")
	}
	if ll.Size != 2 {
		t.Errorf("删除一个节点后，预期链表大小为2，实际为 %d", ll.Size)
	}
	if ll.Contains(2) {
		t.Error("删除后链表不应包含值2")
	}

	// 测试删除不存在的元素
	if ll.Delete(4) {
		t.Error("删除不存在的元素4应返回false")
	}

	// 测试删除头节点
	if !ll.Delete(1) {
		t.Error("删除头节点值1失败")
	}
	if ll.Size != 1 {
		t.Errorf("删除头节点后，预期链表大小为1，实际为 %d", ll.Size)
	}
}
