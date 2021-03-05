package hard

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	num := findMedianSortedArrays([]int{1}, []int{2, 3})
	fmt.Printf("res: %v", num)
}
