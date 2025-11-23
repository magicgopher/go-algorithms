package question03

// lengthOfLongestSubstring 返回 s 中最长不含重复字符的子串的长度
func lengthOfLongestSubstring(s string) int {
	// 记录每个字符最近一次出现的位置
	// 因为题目只包含英文字符、数字、符号、空格，ASCII 码范围 0~127 足够
	lastSeen := make([]int, 128) // 初始化为 0
	for i := range lastSeen {
		lastSeen[i] = -1 // -1 表示该字符尚未出现
	}

	n := len(s)
	maxLen := 0 // 记录全局最大长度
	left := 0   // 滑动窗口的左边界

	for right := 0; right < n; right++ {
		ch := s[right]
		// 如果当前字符在窗口 [left, right) 中已经出现过
		if lastSeen[ch] >= left {
			// 把左边界移动到重复字符出现位置的下一位
			left = lastSeen[ch] + 1
		}
		// 更新该字符的最新位置
		lastSeen[ch] = right

		// 当前窗口长度为 right - left + 1
		curLen := right - left + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
