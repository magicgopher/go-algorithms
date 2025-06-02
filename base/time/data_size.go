package time

// 时间负责度：用来衡量一个算法执行时间随着数据规模增长而变化的度量（定性的描述一个算法的执行时间）
// 数据规模

// Sum 计算从 1 到 n 的所有整数之和，这里入参的 n 就是数据规模
// 这里的n是数值，数据规模也就是int类型数值的大小
func Sum(n int) int {
	res := 0
	i := 0
	for ; i <= n; i++ {
		res += i
	}
	return res
}

// FindMax 在一个整数切片中查找最大值
func FindMax(s []int) int {
	if len(s) == 0 {
		return 0 // 空切片返回0
	}
	m := s[0] // 假设第一个元素就是最大值
	for i := 1; i <= len(s)-1; i++ {
		if s[i] > m {
			m = s[i]
		}
	}
	return m
}

// ListNode 定义单链表节点结构
type ListNode struct {
	Val  int       // Val 存储节点的整数值
	Next *ListNode // Next 指向下一个节点的指针，nil 表示链表末尾
}

// Length 计算单链表的节点数
func Length(head *ListNode) int {
	count := 0           // count 用于记录链表节点数
	current := head      // current 指向当前遍历的节点，初始为头节点
	for current != nil { // 遍历链表，直到遇到 nil（链表末尾）
		count++                // 每访问一个节点，计数加 1
		current = current.Next // 移动到下一个节点
	}
	return count // 返回链表的总节点数
}
