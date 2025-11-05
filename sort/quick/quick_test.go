package quick

import (
	"reflect"
	"testing"
)

// TestQuickSort 测试 QuickSort 函数的各种情况
func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string // 测试名称
		input    []int  // 输入数组
		expected []int  // 期望的排序结果
	}{
		{
			name:     "已排序数组",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序数组",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "无序数组",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "空数组",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单个元素",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "重复元素",
			input:    []int{3, 3, 3, 3},
			expected: []int{3, 3, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入的副本，避免修改原始测试数据
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// 执行快速排序
			QuickSort(input)

			// 比较排序结果是否符合预期
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("QuickSort(%v) = %v; 期望 %v", tt.input, input, tt.expected)
			}
		})
	}
}
