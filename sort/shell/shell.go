package shell

// ShellSort 函数实现希尔排序（就地排序）
// 参数：arr []int - 需要排序的整数切片
// 返回值：无（直接修改原切片）
func ShellSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	n := len(arr)
	// 初始间隔 gap，通常从 n/2 开始，并逐步减半
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对每个子序列进行插入排序
		for i := gap; i < n; i++ {
			// 保存当前元素
			temp := arr[i]
			j := i
			// 在子序列中向前比较并移动元素
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			// 插入到正确位置
			arr[j] = temp
		}
	}
}
