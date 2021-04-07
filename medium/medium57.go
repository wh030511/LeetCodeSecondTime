package medium

import "sort"

// 插入区间
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	le, ri := newInterval[0], newInterval[1]
	for _, value := range intervals {
		l, r := value[0], value[1]
		if r < le {
			res = append(res, value)
			continue
		}
		if l > ri {
			res = append(res, value)
			continue
		}
		if l < le {
			le = l
		}
		if r > ri {
			ri = r
		}
	}
	res = append(res, []int{le, ri})
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}
