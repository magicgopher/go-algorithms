package merge

import (
	"fmt"
	"testing"
)

// TestMergeSort 归并排序单元测试
func TestMergeSort(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2, 9, 1, 7, 6}
	fmt.Println("排序前:", arr)
	sorted := MergeSort(arr)
	fmt.Println("排序后:", sorted)
}
