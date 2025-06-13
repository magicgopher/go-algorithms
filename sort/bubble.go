package sort

// 冒泡排序

func BubbleSort(s []int) []int {
	n := len(s)
	// 创建一个新的切片用于存储排序结果
	result := make([]int, n)
	copy(result, s)

	// 外层循环控制排序轮数
	for i := 0; i < n-1; i++ {
		// 内层循环进行相邻元素比较和交换
		for j := 0; j < n-1-i; j++ {
			if result[j] > result[j+1] {
				// 交换元素
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}
