package question13

import "strings"

// 从左到右遍历 + 判断下一个字符
func romanToInt1(s string) int {
	m := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	ans := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if i < n-1 && m[s[i]] < m[s[i+1]] {
			ans -= m[s[i]] // 例如 I 在 V 左边，要减 1
		} else {
			ans += m[s[i]]
		}
	}
	return ans
}

// 从右到左遍历
func romanToInt2(s string) int {
	m := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	ans := m[s[len(s)-1]] // 先加最后一位
	for i := len(s) - 2; i >= 0; i-- {
		if m[s[i]] < m[s[i+1]] {
			ans -= m[s[i]]
		} else {
			ans += m[s[i]]
		}
	}
	return ans
}

// 替换法
func romanToInt3(s string) int {
	replacements := map[string]string{
		"IV": "A", // A代表4
		"IX": "B", // B代表9
		"XL": "E", // 40
		"XC": "F", // 90
		"CD": "G", // 400
		"CM": "H", // 900
	}

	for old, new := range replacements {
		s = strings.ReplaceAll(s, old, new)
	}

	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
		'A': 4, 'B': 9, 'E': 40, 'F': 90, 'G': 400, 'H': 900}

	ans := 0
	for i := range s {
		ans += m[s[i]]
	}
	return ans
}

// 硬编码所有特殊情况
func romanToInt4(s string) int {
	m := map[string]int{
		"I": 1, "IV": 4, "V": 5, "IX": 9,
		"X": 10, "XL": 40, "L": 50, "XC": 90,
		"C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000,
	}

	ans := 0
	i := 0
	for i < len(s) {
		if i+1 < len(s) && m[s[i:i+2]] > 0 {
			ans += m[s[i:i+2]]
			i += 2
		} else {
			ans += m[s[i:i+1]]
			i++
		}
	}
	return ans
}
