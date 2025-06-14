# Go算法和数据结构

## 简介

这是一个用于记录使用 Go 语言实现的算法和数据结构的仓库。目标是通过实践加深对算法、数据结构以及 Go 编程语言的理解。仓库包含多种经典算法和数据结构的实现，代码清晰且附有注释，适合初学者和进阶开发者学习参考。

## 时间复杂度

时间复杂度（Time Complexity）是用来衡量一个算法执行时间随输入规模增长而变化的度量。它描述了算法运行时间与输入数据规模之间的关系，通常用大O表示法（Big-O Notation）来表示，忽略常数和低阶项，只关注最高阶项的增长趋势。

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

如何理解时间复杂度？实际上，一个算法，在不同的配置的机器上，它的执行时间都是不同的，但是，在不同配置的机器上，这个算法的语句执行次数是一样的【这个算法的语句执行次数是一样的指的是：对于相同的输入数据规模 `n`，算法的逻辑决定了基本操作的执行次数是固定的】，所以呢？可以使用语句的执行次数来衡量算法的运行时间。

从下面这个例子来估算算法的执行时间。

```go
func Sum(n int) int {
	res := 0            // (1) 执行1次
	i := 0              // (2) 执行1次
	for ; i <= n; i++ { // (3) 执行了2n次【i<=n; i++是两个语句，每个语句执行了n次】
		res += i 					// (4) 执行了n次
	}
	return res 					// (5) 执行了1次
}
```

示例分析：

- `res := 0` 这行代码执行了1次。
- `i := 0` 这行代码执行了1次。
- `for` 循环中的 `; i <= n; i++` 其中 `i <= n` 执行了n次，`i++` 也执行了n次，这行共执行了2n次。
- `for` 循环中的 `res += 1` 执行了n次。
- `return res` 执行了1次。
- 总共执行了 3n + 3 次

随着数据规模 `n` 的不断增大，系数和常量对执行时间的增长趋势影响非常小，那么这个 `Sum()` 函数的时间复杂度就是 `O(n)`。

为什么随着数据规模 `n` 的不断增大，系数和常量对执行时间的增长趋势影响非常小呢？

下面来看一组数据，如下：

| n     | n+4   |
| ----- | ----- |
| 1     | 5     |
| 10    | 14    |
| 100   | 104   |
| 1000  | 1004  |
| 10000 | 10004 |

## 空间复杂度

空间复杂度（Space complexity）是指算法在执行过程中所需的额外存储空间（不包括输入数据本身占用的空间）。它衡量算法运行时在内存中使用的空间量，通常用大O表示法来描述，反映随着输入规模 n 的增长，算法所需额外空间的变化趋势【算法的存储空间与数据规模之间的增长关系】。

## 数据结构

- 数组
- 链表
- 堆
- 栈
- 队列
- 哈希表
- 图
- 树

## 算法