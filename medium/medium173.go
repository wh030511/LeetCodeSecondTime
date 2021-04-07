package medium

type BSTIterator struct {
	nums []int
	index int
}

// todo
func Constructor(root *TreeNode) BSTIterator {
	// 前序遍历二叉树，放到数组中
	nums := recur173(root, []int{})
	return BSTIterator{
		nums: nums,
		index: -1,
	}
}

func recur173(node *TreeNode, nums []int) []int {
	if node == nil {
		return nums
	}
	if node.Left == nil && node.Right == nil {
		nums = append(nums, node.Val)
		return nums
	}
	nums = recur173(node.Left, nums)
	nums = append(nums, node.Val)
	nums = recur173(node.Right, nums)
	return nums
}


func (this *BSTIterator) Next() int {
	this.index++
	return this.nums[this.index]
}


func (this *BSTIterator) HasNext() bool {
	return this.index + 1 < len(this.nums)
}

