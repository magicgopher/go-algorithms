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
		table:    make([]*entry, capacity),
		capacity: capacity,
	}
}

// hash 计算字符串 key 的哈希值（用于索引）
func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (31*hash + int(ch)) % ht.capacity
	}
	return hash
}

// Put 向哈希表中插入或更新一个键值对
func (ht *HashTable) Put(key, value string) error {
	if ht.size >= ht.capacity {
		return errors.New("哈希表已满")
	}
	index := ht.hash(key)
	originalIndex := index

	// 使用线性探测处理冲突
	for {
		if ht.table[index] == nil || ht.table[index].key == key {
			// 插入新条目或更新已有键
			ht.table[index] = &entry{key: key, value: value}
			ht.size++
			return nil
		}
		index = (index + 1) % ht.capacity
		if index == originalIndex {
			break
		}
	}
	return errors.New("插入失败，哈希表可能已满")
}

// Get 根据键查找对应的值，如果存在则返回
func (ht *HashTable) Get(key string) (string, bool) {
	index := ht.hash(key)
	originalIndex := index

	for {
		if ht.table[index] == nil {
			return "", false
		}
		if ht.table[index].key == key {
			return ht.table[index].value, true
		}
		index = (index + 1) % ht.capacity
		if index == originalIndex {
			break
		}
	}
	return "", false
}

// Remove 从哈希表中移除一个键值对
func (ht *HashTable) Remove(key string) bool {
	index := ht.hash(key)
	originalIndex := index

	for {
		if ht.table[index] == nil {
			return false
		}
		if ht.table[index].key == key {
			ht.table[index] = nil
			ht.size--
			return true
		}
		index = (index + 1) % ht.capacity
		if index == originalIndex {
			break
		}
	}
	return false
}

// Size 返回当前哈希表中键值对的数量
func (ht *HashTable) Size() int {
	return ht.size
}
