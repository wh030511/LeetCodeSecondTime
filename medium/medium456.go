package medium

import "math"

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	space := make([][]int, 0, len(nums)/2)
	low, high := math.MaxInt32, math.MinInt32
	for _, v := range nums {
		if low < high {
			space = append(space, []int{low, high})
		}
		if low < v && v < high {
			return true
		}
		for _, sp := range space {
			if sp[0] < v && v < sp[1] {
				return true
			}
		}
		if v <= low {
			low = v
			high = math.MinInt32
			continue
		}
		if v > high {
			high = v
			continue
		}
	}
	return false
}
