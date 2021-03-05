package medium

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	arr := []byte(s)
	dic := make(map[byte]int)
	dp := make([]int, len(s))
	dp[0] = 1
	dic[arr[0]] = 0
	for i := 1; i < len(arr); i++ {
		if _, ok := dic[arr[i]]; !ok {
			dp[i] = dp[i-1] + 1
			dic[arr[i]] = i
			continue
		}

		if i-dic[arr[i]] < dp[i-1]+1 {
			dp[i] = i - dic[arr[i]]
		} else {
			dp[i] = dp[i-1] + 1
		}

		dic[arr[i]] = i
	}

	r := 0
	for _, v := range dp {
		if v > r {
			r = v
		}
	}

	return r
}
