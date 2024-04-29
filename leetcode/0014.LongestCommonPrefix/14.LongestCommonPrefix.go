package leetcode

func longestCommonPrefix(strs []string) string {
	resultStr := ""
	ok := true
	l := len(strs[0])

	var currentChar byte
	for i := 0; l > i; i++ {
		currentChar = strs[0][i]

		for _, str := range strs {
			if len(str) <= i {
				ok = false
				break
			}

			if str[i] != currentChar {
				ok = false
				break
			}
		}

		if !ok {
			break
		}

		resultStr += string(currentChar)
	}

	return resultStr
}
