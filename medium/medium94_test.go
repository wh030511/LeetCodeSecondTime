package medium

import (
	"fmt"
	"testing"
)

func TestMedium94(t *testing.T) {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 9,
		},
	}
	res := inorderTraversal(root)
	fmt.Printf("res: %+v", res)
}
