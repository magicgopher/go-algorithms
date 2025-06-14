package bubble

// 冒泡排序
// 属于交换类的排序算法
// 冒泡排序是通过连续地比较与交换相邻元素实现排序
// 过程就像气泡从底部升到顶部一样

func Bubble(s []int) []int {
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
