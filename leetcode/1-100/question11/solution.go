package question11

// 双指针
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// 计算当前面积
		width := right - left
		h := height[left]
		if height[left] > height[right] {
			h = height[right]
			right--
		} else {
			left++
		}
		area := h * width
		if area > maxWater {
			maxWater = area
		}
	}
	return maxWater
}

// 双指针优化写法
func maxArea1(height []int) int {
	left, right := 0, len(height)-1
	ans := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)

		if height[left] < height[right] {
			left++
		} else {
			right-- // 注意：相等时也可以移动右边，效果一样，甚至有时更快
		}
	}
	return ans
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力枚举
func maxArea2(height []int) int {
	n := len(height)
	maxWater := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > maxWater {
				maxWater = area
			}
		}
	}
	return maxWater
}
