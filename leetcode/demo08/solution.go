package demo08

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := 0
	n := len(s)

	// 1. 跳过前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0 // 只有空格
	}

	// 2. 处理符号
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	// 3. 转换数字（跳过前导零，直到非数字停止）
	var result int64 = 0 // 用 int64 防止溢出时计算出错
	for i < n {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		digit := int64(s[i] - '0')
		result = result*10 + digit

		// 4. 溢出判断（提前检测，避免真正溢出后再处理）
		if sign == 1 && result > (1<<31)-1 {
			return 1<<31 - 1 // 返回 2^31 - 1
		}
		if sign == -1 && result > 1<<31 {
			return -1 << 31 // 返回 -2^31
		}

		i++
	}

	return int(result * int64(sign))
}
