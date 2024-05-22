package leetcode

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}

	allPartitions := make([][]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if palindromeSubS, ok := checkSubString(s[:i+1]); ok {
			remaining := partition(s[i+1:])
			for _, r := range remaining {
				partition := append([]string{palindromeSubS}, r...)
				allPartitions = append(allPartitions, partition)
			}
		}
	}

	return allPartitions
}

func checkSubString(s string) (string, bool) {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-(1+i)] {
			return "", false
		}
	}
	return s, true
}
