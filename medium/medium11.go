package medium

func maxArea(height []int) int {
	le, ri := 0, len(height)-1
	res := 0
	for le < ri {
		vo := (ri - le) * minInt(height[le], height[ri])
		if res < vo {
			res = vo
		}

		if height[le] < height[ri] {
			le++
			continue
		}
		ri--
	}
	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
