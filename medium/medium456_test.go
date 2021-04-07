package medium

import (
	"fmt"
	"testing"
)

func TestFind132pattern(t *testing.T) {
	res := find132pattern([]int{3,5,0,3,4})
	fmt.Printf("res: %v", res)
}
