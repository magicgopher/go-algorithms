package demo02

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers 实现两个逆序存储的非负整数链表相加
// 输入: l1, l2 为两个非空链表，节点值 0~9，无前导零
// 输出: 结果链表，同样逆序存储
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建虚拟头节点，避免处理头节点为空的边界情况
	l := &ListNode{}

	// p, q 分别遍历 l1 和 l2，r 是结果链表的当前指针
	p, q, r := l1, l2, l
	carry := 0 // 进位，初始为 0

	// 当还有数字未处理完 或 存在进位 时，继续循环
	for p != nil || q != nil || carry > 0 {
		v := carry // 当前位的和从进位开始

		// 如果 l1 还有节点，累加其值并前移
		if p != nil {
			v += p.Val
			p = p.Next
		}
		// 如果 l2 还有节点，累加其值并前移
		if q != nil {
			v += q.Val
			q = q.Next
		}

		// 计算当前位的值和新的进位
		carry = v / 10 // 进位：十位部分
		r.Val = v % 10 // 当前节点值：个位部分

		// 如果后续还有数字或进位需要处理，则预分配下一个节点
		// 注意：最后一次循环若无进位则不再创建新节点
		if p != nil || q != nil || carry > 0 {
			r.Next = &ListNode{} // 创建新节点
			r = r.Next           // 移动指针到新节点
		}
	}

	// 返回虚拟头节点的下一个节点，即真实结果链表
	return l.Next
}
