package hard

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	t := len(nums1) + len(nums2)
	m := 0
	if t % 2 == 0 {
		m = t / 2 - 1
	} else {
		m = t / 2
	}
	if len(nums1) == 0 {
		return findMid(nums2)
	}
	if len(nums2) == 0 {
		return findMid(nums1)
	}
	p1, p2 := 0, 0

}

func findMid(nums []int) float64 {
	if len(nums) == 0 {
		return float64(0)
	}

	l := len(nums)
	if l % 2 == 0 {
		return float64(nums[l/2 - 1] + nums[l/2]) / 2
	}
	return float64(nums[l/2])
}
