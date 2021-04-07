package medium

import (
	"fmt"
	"testing"
)

func TestMedium173(t *testing.T) {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 5,
			//Left: &TreeNode{
			//	Val: 9,
			//},
			//Right: &TreeNode{
			//	Val: 20,
			//},
		},
	}
	iterator := Constructor(root)
	for iterator.HasNext() {
		fmt.Printf("value: %v", iterator.Next())
	}
}
