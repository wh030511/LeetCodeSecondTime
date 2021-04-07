package medium

func search(nums []int, target int) bool {
	if len(nums) < 2 {
		return binarySearch(nums, target)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return binarySearch(nums[i:], target) || binarySearch(nums[:i], target)
		}
	}
	return binarySearch(nums, target)
}

// 二分查找
func binarySearch(nums []int, target int) bool {
	le, ri := 0, len(nums)
	for le < ri {
		mi := (le+ri) / 2
		if nums[mi] == target {
			return true
		}
		if nums[mi] < target {
			le = mi + 1
		} else {
			ri = mi
		}
	}
	return false
}
