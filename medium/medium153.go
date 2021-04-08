package medium

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mi := (low+high) / 2
		if nums[mi] < nums[high] {
			high = mi
			continue
		}
		if nums[mi] > nums[high] {
			low = mi + 1
			continue
		}
	}
	return nums[low]
}
