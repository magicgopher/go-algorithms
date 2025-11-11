package demo02

import (
	"fmt"
	"testing"
)

// 从切片创建链表，方便写测试用例
func newList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// 把链表转成字符串，方便打印
func listToString(head *ListNode) string {
	s := ""
	for head != nil {
		if s != "" {
			s += " -> "
		}
		s += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}
	return s
}

func TestAddTwoNumbers(t *testing.T) {
	// 示例1
	l1 := newList([]int{2, 4, 3})
	l2 := newList([]int{5, 6, 4})
	fmt.Printf("Example 1: %s + %s = %s\n",
		listToString(l1), listToString(l2), listToString(addTwoNumbers(l1, l2)))

	// 示例2
	l1 = newList([]int{0})
	l2 = newList([]int{0})
	fmt.Printf("Example 2: %s + %s = %s\n",
		listToString(l1), listToString(l2), listToString(addTwoNumbers(l1, l2)))

	// 示例3
	l1 = newList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = newList([]int{9, 9, 9, 9})
	fmt.Printf("Example 3: %s + %s = %s\n",
		listToString(l1), listToString(l2), listToString(addTwoNumbers(l1, l2)))
}
