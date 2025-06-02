package time

import (
	"fmt"
	"testing"
)

// TestSum Sum单元测试
func TestSum(t *testing.T) {
	res := Sum(5)
	fmt.Println(res)
}

func TestFindMax(t *testing.T) {
	nums := []int{18, 22, 6, 29, 51}
	res := FindMax(nums)
	fmt.Println(res)
}

func TestListNodeLength(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name     string    // 测试用例名称
		head     *ListNode // 输入链表的头节点
		expected int       // 期望的链表长度
	}{
		{
			name:     "Empty list", // 测试空链表
			head:     nil,
			expected: 0,
		},
		{
			name:     "Single node", // 测试单节点链表
			head:     &ListNode{Val: 1, Next: nil},
			expected: 1,
		},
		{
			name:     "Multiple nodes", // 测试多节点链表
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
			expected: 3,
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 调用 Length 函数并获取实际结果
			result := Length(tt.head)
			// 比较实际结果与期望结果
			if result != tt.expected {
				t.Errorf("Length(%v) = %d; want %d", tt.head, result, tt.expected)
			}
		})
	}
}
