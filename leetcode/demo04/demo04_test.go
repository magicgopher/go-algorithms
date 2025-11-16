package demo04

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{1, 3}, []int{2}))       // 2.0
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))    // 2.5
	t.Log(findMedianSortedArrays([]int{}, []int{1}))           // 1.0
	t.Log(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5})) // 3.0
}
