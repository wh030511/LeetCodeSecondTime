package medium

func longestPalindrome(s string) string {
	bytes := []byte(s)
	res := ""
	for index, _ := range bytes {
		r := helper(bytes, index)
		if len(r) > len(res) {
			res = r
		}
	}
	return res
}

func helper(str []byte, index int) string {
	l, r := index-1, index+1
	ret := string(str[index : index+1])

	var ret2 string
	if r < len(str) && str[index] == str[r] { // 特殊情况
		ret2 = helper2(str, index)
	}

	for l >= 0 && r < len(str) && str[l] == str[r] {
		ret = string(str[l : r+1])
		l--
		r++
	}

	if len(ret) < len(ret2) {
		return ret2
	}

	return ret
}

func helper2(str []byte, index int) string {
	l, r := index-1, index+2
	ret := string(str[index : index+2])

	for l >= 0 && r < len(str) && str[l] == str[r] {
		ret = string(str[l : r+1])
		l--
		r++
	}

	return ret
}
