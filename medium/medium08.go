package medium

import (
	"math"
	"strconv"
)

func myAtoi(s string) int {
	str := []byte(s)
	l := 0
	for index := range str {
		if str[index] != byte(' ') {
			break
		}
		l++
	}
	str = str[l:]

	if len(str) == 0 {
		return 0
	}

	if str[0] == '+' || str[0] == '-' {
		res := []byte{str[0]}
		for i := 1; i < len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				break
			}
			res = append(res, str[i])
		}
		num, _ := strconv.ParseInt(string(res), 10, 64)
		if num < math.MinInt32 {
			return math.MinInt32
		}
		if num > math.MaxInt32 {
			return math.MaxInt32
		}
		return int(num)
	}

	if str[0] < '0' || str[0] > '9' {
		return 0
	}

	res := []byte{}
	for i := range str {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		res = append(res, str[i])
	}
	num, _ := strconv.ParseInt(string(res), 10, 64)
	if num < math.MinInt32 {
		return math.MinInt32
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(num)
}
