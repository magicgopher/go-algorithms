package heap

// 常用数据结构
// 堆（Heap）

// Heap 表示一个最小堆的结构体
type Heap struct {
	items []int // 使用int类型切片来保存元素
}

// New 创建一个空的最小堆
func New() *Heap {
	return &Heap{items: []int{}}
}

// Push 插入一个值，保持最小堆性质
func (h *Heap) Push(value int) {
	h.items = append(h.items, value)
	h.siftUp(len(h.items) - 1)
}

// Pop 删除并返回堆顶（最小值）
// 如果堆为空，返回 -1 和 false
func (h *Heap) Pop() (int, bool) {
	if h.IsEmpty() {
		return -1, false
	}
	min := h.items[0]
	last := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	if len(h.items) > 0 {
		h.items[0] = last
		h.siftDown(0)
	}
	return min, true
}

// Peek 查看堆顶（最小值）但不删除
// 如果堆为空，返回 -1 和 false
func (h *Heap) Peek() (int, bool) {
	if h.IsEmpty() {
		// 堆为空的情况
		return -1, false
	}
	// 堆不为空的情况
	return h.items[0], true
}

// Size 返回堆中的元素数量
func (h *Heap) Size() int {
	return len(h.items)
}

// IsEmpty 检查堆是否为空
func (h *Heap) IsEmpty() bool {
	return len(h.items) == 0
}

// siftUp 上浮操作，确保插入后保持最小堆性质
func (h *Heap) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[index] >= h.items[parent] {
			break
		}
		h.items[index], h.items[parent] = h.items[parent], h.items[index]
		index = parent
	}
}

// siftDown 下沉操作，确保删除后保持最小堆性质
func (h *Heap) siftDown(index int) {
	minIndex := index
	for {
		left := 2*index + 1
		right := 2*index + 2
		if left < len(h.items) && h.items[left] < h.items[minIndex] {
			minIndex = left
		}
		if right < len(h.items) && h.items[right] < h.items[minIndex] {
			minIndex = right
		}
		if minIndex == index {
			break
		}
		h.items[index], h.items[minIndex] = h.items[minIndex], h.items[index]
		index = minIndex
	}
}
