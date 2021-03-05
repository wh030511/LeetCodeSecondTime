package easy

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numss := [][]int{
		{2,7,11,5},
		{3,2,4},
		{3,3},
	}
	targets := []int{9, 6, 6}

	for i := 0; i < 3; i++ {
		res := twoSum(numss[i], targets[i])
		fmt.Printf("nums: %+v, target: %d, res: %+v \n", numss[i], targets[i], res)
	}
}
