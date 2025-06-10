package array

// 数据结构
// 动态数组

type Array struct {
	items []int // 动态数组存储的元素是int类型
}

// New 创建一个Array指针
func New() *Array {
	initArray := make([]int, 0, 10) // 创建初始化切片长度为0，容量为10
	return &Array{items: initArray}
}

// Append 在动态数组末尾追加一个元素
func (a *Array) Append(item int) {
	a.items = append(a.items, item)
}

// Get 获取指定索引处的元素
func (a *Array) Get(index int) (int, bool) {
	if index < 0 || index >= len(a.items) {
		// 无效索引
		return -1, false
	}
	return a.items[index], true
}

// Set 将指定索引处的元素设置为给定值
func (a *Array) Set(index int, item int) bool {
	if index < 0 || index >= len(a.items) {
		return false
	}
	a.items[index] = item // 设置新的值
	return true
}

// Remove 删除指定索引处的元素，并返回该元素的值
func (a *Array) Remove(index int) (int, bool) {
	if index < 0 || index >= len(a.items) {
		return 0, false
	}
	item := a.items[index]
	a.items = append(a.items[:index], a.items[index+1:]...)
	return item, true
}

// Size 返回数组中的元素数量。
func (a *Array) Size() int {
	return len(a.items)
}

// IsEmpty 检查数组是否为空。
func (a *Array) IsEmpty() bool {
	return len(a.items) == 0
}
