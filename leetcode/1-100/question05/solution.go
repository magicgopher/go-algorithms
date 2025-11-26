package question05

// 暴力查找
func longestPalindrome1(s string) string {
	if len(s) == 1 {
		return s
	}
	ans := ""
	ansLength := 0

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i+ansLength; j-- {
			subString := s[i:j]
			if isPalindrome(subString) && len(subString) > ansLength {
				ans = subString
				ansLength = len(subString)
			}
		}
	}

	return ans
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

// 马拉车
func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}

	// 预处理：在字符间插入 '#'，避免奇偶长度回文单独处理
	// 例如 "aba" -> "#a#b#a#"
	t := "#"
	for _, ch := range s {
		t += string(ch) + "#"
	}

	n := len(t)
	// P[i] 表示以 t[i] 为中心的最长回文半径（包含中心）
	P := make([]int, n)
	center, right := 0, 0 // 当前能扩展到的最右边界和对应的中心

	maxLen, maxCenter := 0, 0

	for i := 0; i < n; i++ {
		// 如果 i 在 right 范围内，可以利用对称性初始化 P[i]
		if i < right {
			mirror := 2*center - i
			P[i] = min(right-i, P[mirror])
		} else {
			P[i] = 0
		}

		// 尝试向两边扩展
		left, rightPos := i-(1+P[i]), i+(1+P[i])
		for left >= 0 && rightPos < n && t[left] == t[rightPos] {
			P[i]++
			left--
			rightPos++
		}

		// 更新最右边界
		if i+P[i] > right {
			center = i
			right = i + P[i]
		}

		// 更新答案
		if P[i] > maxLen {
			maxLen = P[i]
			maxCenter = i
		}
	}

	// 从预处理字符串 t 中提取原字符串的最长回文子串
	start := (maxCenter - maxLen) / 2 // 在原字符串中的起始位置
	return s[start : start+maxLen]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func longestPalindrome3(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				start = left
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		// 奇数长度回文，以 i 为中心
		expand(i, i)
		// 偶数长度回文，以 i 和 i+1 为中心
		expand(i, i+1)
	}

	return s[start : start+maxLen]
}
