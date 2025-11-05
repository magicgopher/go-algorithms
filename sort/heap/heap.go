package heap

// HeapSort 使用堆排序算法对切片进行升序排序。
// 它会直接在原输入切片上进行修改（原地排序）。
func HeapSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	// 步骤1：从最后一个非叶子节点开始，构建最大堆
	// 最后一个非叶子节点的索引是 (n/2 - 1)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 步骤2：逐个将堆顶（最大值）放到数组末尾，并重新调整堆
	for i := n - 1; i > 0; i-- {
		// 将当前最大值（堆顶）交换到末尾
		arr[0], arr[i] = arr[i], arr[0]
		// 对剩余未排序部分（0 到 i-1）重新调整为最大堆
		heapify(arr, i, 0)
	}
}

// heapify 维护最大堆性质：父节点 >= 左右子节点
// n 是当前堆的大小，i 是需要下沉的节点索引
func heapify(arr []int, n, i int) {
	largest := i     // 暂定当前节点为最大
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点存在且大于当前最大值
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在且大于当前最大值
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是当前节点，则交换并继续下沉
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// 递归调整受影响的子树
		heapify(arr, n, largest)
	}
}
