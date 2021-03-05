package medium

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		twoRet := twoSum(nums[i+1:], 0-v)
		for _, slices := range twoRet {
			temp := make([]int, 0)
			temp = append(temp, v)
			temp = append(temp, slices...)
			res = append(res, temp)
		}
	}

	return res
}

func twoSum(nums []int, target int) [][]int {
	le, ri := 0, len(nums)-1
	res := make([][]int, 0)

	for le < ri {
		if le > 0 && nums[le-1] == nums[le] {
			le++
			continue
		}
		if ri < len(nums)-1 && nums[ri+1] == nums[ri] {
			ri--
			continue
		}
		if nums[le]+nums[ri] == target {
			res = append(res, []int{nums[le], nums[ri]})
			le++
			ri--
			continue
		}
		if nums[le]+nums[ri] < target {
			le++
			continue
		}
		if nums[le]+nums[ri] > target {
			ri--
			continue
		}
	}
	return res
}
