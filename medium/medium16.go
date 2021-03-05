package medium

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := nums[0] + nums[1] + nums[2]
	for i, v := range nums {
		if i >= len(nums)-2 {
			break
		}
		r := helper16(nums[i+1:], target-v)
		if calculateS(r+v, target) < calculateS(res, target) {
			res = r + v
		}
		if res == target {
			return res
		}
	}
	return res
}

func helper16(nums []int, target int) int {
	if len(nums) < 2 {
		return 0
	}
	le, ri := 0, len(nums)-1
	res := nums[le] + nums[ri]

	for le < ri {
		value := nums[le] + nums[ri]
		if calculateS(value, target) < calculateS(res, target) {
			res = value
		}
		if value == target {
			return value
		}
		if value < target {
			le++
			continue
		}
		if value > target {
			ri--
			continue
		}
	}

	return res
}

func calculateS(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
