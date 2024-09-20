package leetcode

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	if palindrome, ok := checkSubString(s); ok {
		return palindrome
	}

	res := ""
	for i := 1; i < len(s); i++ {
		res = res + string(s[len(s)-i])
		if palindrome, ok := checkSubString(res + s); ok {
			return palindrome
		}
	}

	return ""
}

func checkSubString(s string) (string, bool) {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-(1+i)] {
			return "", false
		}
	}
	return s, true
}
