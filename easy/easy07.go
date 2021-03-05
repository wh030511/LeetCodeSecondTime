package easy

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	flag := false
	if x < 0 {
		flag = true
		x = -x
	}
	str := []byte(strconv.Itoa(x))
	l, r := 0, len(str)-1
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
	num, _ := strconv.ParseInt(string(str), 10, 64)
	if flag {
		num = -num
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return int(num)
}
