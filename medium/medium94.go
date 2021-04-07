package medium

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0, 0)
	recur94(root, nums)
	return nums
}

func recur94(node *TreeNode, nums []int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		nums = append(nums, node.Val)
		return
	}
	recur94(node.Left, nums)
	nums = append(nums, node.Val)
	recur94(node.Right, nums)
}
