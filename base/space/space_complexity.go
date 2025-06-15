package space

// 空间复杂度是指算法在运行的过程中所需要的额外内存空间大小，通常用大O表示法描述。
// 它包括：固定部分、可变部分
// 固定部分：如变量、常量等，与输入数据规模无关。
// 可变部分：如递归调用栈、动态分配的数组等，与输入规模相关。

// SumArray 函数的空间复杂度为O(1)
// 分析：无论 arr 有多大，只用 total 一个变量，空间复杂度为 O(1)
func SumArray(arr []int) int {
	total := 0 // 定义一个变量存储和
	for _, v := range arr {
		total += v
	}
	return total
}

// binarySearch 函数的空间复杂度为O(log n)
// 常见于分治算法（如二分查找或某些递归算法），每次递归将问题规模减半
// 分析：递归调用栈深度为 log n（每次问题规模减半），空间复杂度为 O(log n)
func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	}
	return binarySearch(arr, target, mid+1, right)
}

// fibonacci 斐波那契
// 函数的空间复杂度为O(n)
// 分析：递归调用栈深度为 n，空间复杂度为 O(n)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// CreateMatrix 常见于需要分配二维数组或类似结构的算法
// 分析：分配一个 n×n 的二维数组，额外空间为 n * n，空间复杂度为 O(n²)
func CreateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

// Subsets 指数增长的空间复杂度通常出现在递归算法中，每次调用产生多个分支，导致调用栈或数据结构呈指数级增长。
// 以下是一个解决子集问题的算法示例，生成给定数组的所有子集。
// 分析：
// 该算法生成数组 nums 的所有子集。对于长度为 n 的数组，子集总数为 2ⁿ
// result 存储所有子集，每个子集的长度最多为 n，总空间需求为 O(2ⁿ * n)（子集数量 × 每个子集的长度）
// 忽略子集长度 n（在复杂度分析中常忽略较低阶项），空间复杂度简化为 O(2ⁿ)
// 注意：递归版本的子集生成（如通过递归调用分支）也可能导致 O(n) 的调用栈深度，但迭代版本（如上）直接存储所有子集，空间需求仍为 O(2ⁿ)
func Subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		newSubsets := make([][]int, 0)
		for _, subset := range result {
			newSubset := make([]int, len(subset))
			copy(newSubset, subset)
			newSubset = append(newSubset, num)
			newSubsets = append(newSubsets, newSubset)
		}
		result = append(result, newSubsets...)
	}
	return result
}
