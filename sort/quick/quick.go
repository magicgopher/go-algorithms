package quick

// QuickSort 使用快速排序算法对切片进行升序排序。
// 它会直接在原输入切片上进行修改（原地排序）。
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

// quickSort 是递归调用的辅助函数。
func quickSort(arr []int, low, high int) {
	if low < high {
		// 分区操作，返回基准元素的最终位置
		pivotIndex := partition(arr, low, high)
		// 递归排序左半部分
		quickSort(arr, low, pivotIndex-1)
		// 递归排序右半部分
		quickSort(arr, pivotIndex+1, high)
	}
}

// partition 选择一个基准值，将数组分区：
// 小于等于基准值的元素放在左边，大于基准值的元素放在右边。
func partition(arr []int, low, high int) int {
	// 选择最右边的元素作为基准值
	pivot := arr[high]

	// i 表示小于等于基准值的区域的右边界（初始为 low-1）
	i := low - 1

	// 从 low 到 high-1 遍历
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++                             // 扩展小于等于区域
			arr[i], arr[j] = arr[j], arr[i] // 交换元素
		}
	}

	// 将基准值放到正确的位置（即 i+1）
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // 返回基准值的最终位置
}
