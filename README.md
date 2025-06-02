# Go算法和数据结构

## 简介

这是一个用于记录使用 Go 语言实现的算法和数据结构的仓库。目标是通过实践加深对算法、数据结构以及 Go 编程语言的理解。仓库包含多种经典算法和数据结构的实现，代码清晰且附有注释，适合初学者和进阶开发者学习参考。

## 时间复杂度

时间复杂度是用来衡量一个算法执行时间随输入规模增长而变化的度量。它描述了算法运行时间与输入数据规模之间的关系，通常用大O表示法（Big-O Notation）来表示，忽略常数和低阶项，只关注最高阶项的增长趋势。

数据规模：一般体现在算法的输入参数中。

- 数值：数据规模就是数值的大小。

```go
// Sum 计算从 1 到 n 的所有整数之和，这里入参的 n 就是数据规模
func Sum(n int) int {
	res := 0
	i := 0
	for ; i <= n; i++ {
		res += i
	}
	return res
}
```

- 数组[切片]：数据规模就是数组的长度。

```go
// FindMax 在一个整数切片中查找最大值，这里的入参的 n 就是数据规模
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
```

- 链表：数据规模就是链表的节点数量。

```go
// ListNode 链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// Length 计算单链表的长度
func Length(head *ListNode) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
```

如何理解时间复杂度？实际上，一个算法，在不同的配置的机器上，它的执行时间都是不同的，但是，在不同配置的机器上，这个算法的语句执行次数是一样的。

## 空间复杂度

## 数据结构

- 数组（Array）
- 

## 算法