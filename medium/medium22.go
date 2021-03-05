package medium

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var res, pre []string
	for ; n > 0; n-- {
		pre = res
		res = nil
		for _, v := range pre {
			ret := helper22(v)
			res = append(res, ret...)
		}
		var last []byte
		for i := n; i > 0; i-- {
			last = append(last, []byte{'(', ')'}...)
		}
		res = append(res, string(last))
		pre = nil
	}

	return res
}

func helper22(s string) []string {
	var res []string
	bytes := []byte(s)
	for i, v := range bytes {
		if v == ')' {
			fir := bytes[:i]
			sec := bytes[i:]
			var temp []byte
			temp = append(temp, fir...)
			temp = append(temp, []byte{'(',')'}...)
			temp = append(temp, sec...)
			res = append(res, string(temp))
		}
	}
	return res
}