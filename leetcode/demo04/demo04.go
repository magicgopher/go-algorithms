package demo04

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	// 保证 nums1 是较短的数组，减少二分空间
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	imin, imax := 0, m
	halfLen := (m + n + 1) / 2 // 左半部分应有的元素个数（向上取整）

	for imin <= imax {
		i := (imin + imax) / 2 // nums1 的划分点（左边有 i 个元素）
		j := halfLen - i       // nums2 的划分点

		if i < m && nums2[j-1] > nums1[i] {
			// i 太小了，需要右移
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i 太大了，需要左移
			imax = i - 1
		} else {
			// 找到有效划分
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft) // 奇数个元素，中位数就是左半部分最大值
			}

			var minRight int
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}

	return 0.0 // 理论上不会执行到这里
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
