package medium

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		if i >= len(nums) - 3 {
			break
		}
		thrRet := thrSum(nums[i+1:], target-v)
		for _, value := range thrRet {
			temp := make([]int, 0)
			temp = append(temp, v)
			temp = append(temp, value...)
			res = append(res, temp)
		}
	}
	return res
}

func thrSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		if i >= len(nums) - 2 {
			break
		}
		twoRet := twoSum2(nums[i+1:], target-v)
		for _, value := range twoRet {
			temp := make([]int, 0)
			temp = append(temp, v)
			temp = append(temp, value...)
			res = append(res, temp)
		}
	}
	return res
}

func twoSum2(nums []int, target int) [][]int {
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
