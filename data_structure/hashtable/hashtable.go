package hashtable

import "errors"

// 数据结构
// 哈希表

// entry 表示哈希表中的一个键值对（kv都是字符串类型）
type entry struct {
	key   string // 键
	value string // 值
}

// HashTable 哈希表结构体
type HashTable struct {
	table    []*entry // 存储键值对的哈希表
	capacity int      // 哈希表容量
	size     int      // 当前哈希表键值对数量
}

// New 创建一个新的哈希表
func New(capacity int) *HashTable {
	return &HashTable{
		table:    make([]*entry, capacity), // 初始化哈希表为指定大小的空桶数组
		capacity: capacity,
	}
}

// hash 计算字符串 key 的哈希值（用于索引）
func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (31*hash + int(ch)) % ht.capacity // 使用一个常见的31基字符串哈希算法
	}
	return hash
}

// Put 向哈希表中插入或更新一个键值对
func (ht *HashTable) Put(key, value string) error {
	// 如果当前哈希表满了，拒绝插入新的键值对
	if ht.size >= ht.capacity {
		return errors.New("哈希表已满")
	}

	index := ht.hash(key)  // 计算初始哈希索引
	originalIndex := index // 记录原始索引，用于判断是否回到了起点

	// 开始线性探测，查找空位或已有的 key
	for {
		if ht.table[index] == nil {
			// 如果当前位置为空，直接插入新 entry
			ht.table[index] = &entry{key: key, value: value}
			ht.size++
			return nil
		} else if ht.table[index].key == key {
			// 如果 key 已存在，更新其对应的 value
			ht.table[index].value = value
			return nil
		}
		// 向后探测下一个索引位置（线性探测）
		index = (index + 1) % ht.capacity

		// 如果回到原始位置，说明哈希表已经探测一圈
		if index == originalIndex {
			break
		}
	}
	return errors.New("插入失败，哈希表可能已满")
}

// Get 根据键查找对应的值，如果存在则返回
func (ht *HashTable) Get(key string) (string, bool) {
	index := ht.hash(key)  // 计算初始哈希索引
	originalIndex := index // 保存原始索引以防止死循环

	// 线性探测查找对应的 key
	for {
		if ht.table[index] == nil {
			// 如果当前位置为空，则说明 key 不存在
			return "", false
		}
		if ht.table[index].key == key {
			// 如果找到匹配的 key，返回对应的 value
			return ht.table[index].value, true
		}
		// 向后继续探测
		index = (index + 1) % ht.capacity

		// 如果已经回到起点，终止查找
		if index == originalIndex {
			break
		}
	}
	return "", false // 未找到 key
}

// Remove 从哈希表中移除一个键值对
func (ht *HashTable) Remove(key string) bool {
	index := ht.hash(key)  // 计算哈希索引
	originalIndex := index // 记录起始索引

	// 线性探测查找要删除的 key
	for {
		if ht.table[index] == nil {
			// 找到空位说明 key 不存在
			return false
		}
		if ht.table[index].key == key {
			// 找到 key，直接将该位置设为 nil（此实现中不处理删除带来的探测链断裂）
			ht.table[index] = nil
			ht.size--
			return true
		}
		index = (index + 1) % ht.capacity

		if index == originalIndex {
			break
		}
	}
	return false // 未找到 key
}

// Size 返回当前哈希表中键值对的数量
func (ht *HashTable) Size() int {
	return ht.size
}
