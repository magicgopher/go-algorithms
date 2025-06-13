package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := []int{66, 34, 25, 12, 22, 11, 90}
	fmt.Println("排序前:", s)
	sort := BubbleSort(s)
	fmt.Println("排序后:", sort)
}
