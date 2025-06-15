package selection

// 选择排序
// 工作原理：在未排序序列中找到最小（或最大）元素
// 存放到排序序列的起始位置，然后再从剩余未排序元素中继续寻找最小（或最大）元素
// 然后放到已排序序列的末尾，以此类推，直到所有元素均排序完毕

// Selection 选择排序
func Selection(nums []int) {
	n := len(nums)
	// 外层循环遍历整个数组
	for i := 0; i < n-1; i++ {
		// 假设当前索引 i 的元素是未排序部分的最小值
		minIndex := i

		// 内层循环从未排序部分的第二个元素开始，寻找真正的最小值
		for j := i + 1; j < n; j++ {
			// 如果发现更小的元素，则更新最小值的索引
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		// 将找到的最小值与未排序部分的第一个元素交换位置
		// 如果 minIndex 没有改变，说明当前元素已经是最小值，无需交换
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}
