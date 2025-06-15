package bubble

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{66, 34, 25, 12, 22, 11, 90}
	fmt.Println("排序前:", s)
	sort := Bubble(s)
	fmt.Println("排序后:", sort)
}
