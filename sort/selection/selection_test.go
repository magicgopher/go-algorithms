package selection

import (
	"fmt"
	"testing"
)

func TestSelection(t *testing.T) {
	data := []int{24, 33, 6, 82, 17, 59, 99, 101, 45}
	fmt.Println("排序前：", data)
	Selection(data)
	fmt.Println("排序后：", data)
}
