package leetcode

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if subS, ok := checkSubString(s[i:j]); ok {
				if len(longest) < len(subS) {
					longest = subS
				}
			}
		}
	}
	return longest
}

func checkSubString(s string) (string, bool) {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-(1+i)] {
			return "", false
		}
	}
	return s, true
}
