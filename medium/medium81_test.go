package medium

import (
	"fmt"
	"testing"
)

func TestMedium81(t *testing.T) {
	res := search([]int{2, 5, 6, 0, 0, 1, 2}, 0)
	fmt.Printf("res: %v", res)
}
