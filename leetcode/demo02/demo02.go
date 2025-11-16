package demo02

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers 实现两个逆序存储的非负整数链表相加
// 输入: l1, l2 为两个非空链表，节点值 0~9，无前导零
// 输出: 结果链表，同样逆序存储
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头节点，便于处理结果链表
	dummy := &ListNode{}
	curr := dummy
	carry := 0 // 进位

	// 遍历两个链表，直到都处理完且无进位
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // 初始为进位

		// 加上 l1 的当前值（如果存在）
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// 加上 l2 的当前值（如果存在）
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 计算当前位和新的进位
		carry = sum / 10
		digit := sum % 10

		// 创建新节点
		curr.Next = &ListNode{Val: digit}
		curr = curr.Next
	}

	return dummy.Next
}
