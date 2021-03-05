package medium

import (
	"fmt"
	"testing"
)

func TestMedium18(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Printf("res: %v", res)
}
