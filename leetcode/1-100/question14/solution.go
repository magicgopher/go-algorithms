package question14

import "strings"

// 横向扫描
// 思路：先取第一个字符串为基准，然后依次跟后面的每个字符串求公共前缀，逐步缩短。
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// 纵向扫描（逐个字符比较）
// 思路：从第0列开始比较所有字符串的同一个位置的字符，只要有一行不一样或某行结束就停止。
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 分治法
// 思路：把数组分成两半，递归求左右两边的最长公共前缀，再合并。
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return helper(strs, 0, len(strs)-1)
}

func helper(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}
	mid := (l + r) / 2
	left := helper(strs, l, mid)
	right := helper(strs, mid+1, r)
	return commonPrefix(left, right)
}

func commonPrefix(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a[:minLen]
}
