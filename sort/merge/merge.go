package merge

// MergeSort 对整数切片进行归并排序
// 返回一个新的有序切片（不修改原切片）
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 分割数组
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// 合并已排序的左右部分
	return merge(left, right)
}

// merge 合并两个已排序的切片
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// 比较并合并
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 添加剩余元素
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
