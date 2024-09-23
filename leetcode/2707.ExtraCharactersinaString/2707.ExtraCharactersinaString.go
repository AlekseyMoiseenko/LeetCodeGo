package leetcode

func minExtraChar(s string, dictionary []string) int {
	wordMap := make(map[string]bool)
	for _, word := range dictionary {
		wordMap[word] = true
	}

	l := len(s)
	dp := make([]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = l
	}
	dp[0] = 0

	for end := 1; end <= l; end++ {
		for start := 0; start < end; start++ {
			sub := s[start:end]
			if wordMap[sub] {
				dp[end] = min(dp[end], dp[start])
			}
		}
		dp[end] = min(dp[end], dp[end-1]+1)
	}

	return dp[l]
}
