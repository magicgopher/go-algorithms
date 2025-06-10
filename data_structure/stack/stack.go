package stack

// 常见数据结构
// 栈
// 特点：先进后出

// Stack 栈（LIFO）
type Stack struct {
	items []int // 栈中的元素，数据类型是int类型
}

// New 创建一个新的栈指针
func New() *Stack {
	return &Stack{items: []int{}} // 栈的items是一个空的切片
}

// Push 将一个元素添加到栈中
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop 移除栈顶的元素，并且返回移除的元素值和结果
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		// 栈为空的情况
		return 0, false
	}
	item := s.items[len(s.items)-1]    // 取出栈中最后的一个元素
	s.items = s.items[:len(s.items)-1] // 将栈中剩余的元素保留
	return item, true
}

// IsEmpty 检查栈中是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Peek 返回栈顶的元素，但不移除
func (s Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		// 栈为空的情况
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Size 获取栈中的元素数量
func (s *Stack) Size() int {
	return len(s.items)
}
