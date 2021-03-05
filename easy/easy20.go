package easy

func isValid(s string) bool {
	dic := make([]byte, 0)
	for _, v := range []byte(s) {
		if v == '(' || v == '[' || v == '{' {
			dic = append(dic, v)
			continue
		}
		if len(dic) == 0 {
			return false
		}
		if v == ')' {
			if dic[len(dic)-1] != '(' {
				return false
			}
		}
		if v == ']' {
			if dic[len(dic)-1] != '[' {
				return false
			}
		}
		if v == '}' {
			if dic[len(dic)-1] != '{' {
				return false
			}
		}
		dic = dic[:len(dic)-1]
	}
	return len(dic) == 0
}
