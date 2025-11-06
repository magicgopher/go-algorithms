package shell

import (
	"fmt"
	"testing"
)

// TestShellSort 希尔排序单元测试
func TestShellSort(t *testing.T) {
	// 测试示例
	arr := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 3}
	fmt.Println("排序前:", arr)
	ShellSort(arr)
	fmt.Println("排序后:", arr) // 输出: 排序后: [3 11 12 22 25 34 45 64 88 90]
}
