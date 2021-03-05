package medium

import (
	"fmt"
	"testing"
)

func TestMedium16(t *testing.T) {
	res := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	fmt.Printf("res: %v", res)
}
