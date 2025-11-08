package demo01

// twoSum 返回两个数的下标，使得它们的和等于 target
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func twoSum(nums []int, target int) []int {
	// 哈希表：value -> index
	hashMap := make(map[int]int)

	for i, val := range nums {
		comp := target - val

		// 如果补数已在哈希表中，说明找到答案
		if idx, isFound := hashMap[comp]; isFound {
			return []int{i, idx}
		}

		// 否则将当前数加入哈希表
		hashMap[val] = i
	}

	// 题目保证有解，所以不会走到这里
	return []int{-1, -1}
}
