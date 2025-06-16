package insertion

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	data := []int{22, 13, 99, 5, 87, 55, 31}
	fmt.Println("排序前:", data)
	Insertion(data)
	fmt.Println("排序后:", data)
}
