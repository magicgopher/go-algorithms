package linked_list

// 常见的数据结构
// 链表

// Node 链表节点
type Node struct {
	Data int   // 链表节点存储的数据
	Next *Node // 指向下一个节点的指针
}

// LinkedList 链表结构
type LinkedList struct {
	Head *Node // 链表头节点指针
	Size int   // 链表节点数量
}

// NewLinkedList 创建新链表
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

// Insert 在链表尾部插入新节点
func (ll *LinkedList) Insert(data int) {
	newNode := &Node{Data: data} // 创建新节点
	if ll.Head == nil {
		// 如果链表为空，新节点成为头节点
		ll.Head = newNode
	} else {
		// 遍历到链表尾部
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		// 将新节点链接到尾部
		current.Next = newNode
	}
	ll.Size++ // 更新链表大小
}

// Delete 删除链表中第一个值为data的节点
func (ll *LinkedList) Delete(data int) bool {
	if ll.Head == nil {
		// 空链表不能删除
		return false
	}

	if ll.Head.Data == data {
		// 如果头节点是要删除的节点
		ll.Head = ll.Head.Next // 更新头节点
		ll.Size--              // 减少链表大小
		return true
	}

	// 遍历链表寻找要删除的节点
	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			// 找到目标节点，跳过它
			current.Next = current.Next.Next
			ll.Size-- // 减少链表大小
			return true
		}
		current = current.Next
	}
	// 链表中未找到要删除的节点
	return false
}

// Contains 检查链表是否包含某个值
func (ll *LinkedList) Contains(data int) bool {
	current := ll.Head
	for current != nil {
		if current.Data == data {
			// 找到目标值
			return true
		}
		current = current.Next // 移动到下一个节点
	}
	// 未找到目标值
	return false
}
