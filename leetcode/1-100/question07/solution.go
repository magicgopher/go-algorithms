package question07

import "math"

func reverse1(x int) int {
	var result int
	for x != 0 {
		// 取最后一位数字
		pop := x % 10
		x /= 10

		// 正数溢出检查：result > MaxInt32/10 或 result == MaxInt32/10 且 pop > 7
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		// 负数溢出检查：result < MinInt32/10 或 result == MinInt32/10 且 pop < -8
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && pop < -8) {
			return 0
		}

		result = result*10 + pop
	}

	return result
}

func reverse2(x int) int {
	result := 0
	for ; x != 0; x /= 10 {
		result = result*10 + x%10
	}

	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}
