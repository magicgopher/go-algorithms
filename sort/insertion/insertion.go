package insertion

// 插入排序

func Insertion(nums []int) {
	n := len(nums)
	// 从第二个元素开始遍历，因为第一个元素本身可以看作是已排序的
	for i := 1; i < n; i++ {
		// key 是当前需要被插入的元素
		key := nums[i]
		// j 是已排序区域的最后一个元素的索引
		j := i - 1

		// 从后向前扫描已排序区域，为 key 找到插入位置
		// 如果已排序的元素大于 key，则将该元素向右移动
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		// 将 key 插入到正确的位置
		nums[j+1] = key
	}
}
