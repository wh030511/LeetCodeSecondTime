package medium

func letterCombinations(digits string) []string {
	dic := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	res := make([]string, 0)
	pre := make([]string, 0)
	for _, v := range []byte(digits) {
		index := int(v) - 48
		if len(res) == 0 {
			res = append(res, dic[index]...)
			continue
		}
		pre = res
		res = nil
		for _, c := range dic[index] {
			for _, existed := range pre {
				temp := ""
				temp += existed + c
				res = append(res, temp)
			}
		}
	}
	return res
}
