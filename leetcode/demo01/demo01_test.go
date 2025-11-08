package demo01

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// 示例 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("输入: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("输出: %v\n\n", twoSum(nums1, target1)) // [0 1]

	// 示例 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Printf("输入: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("输出: %v\n\n", twoSum(nums2, target2)) // [1 2]

	// 示例 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Printf("输入: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("输出: %v\n", twoSum(nums3, target3)) // [0 1]
}
